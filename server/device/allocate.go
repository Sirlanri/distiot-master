package device

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

/* 传入dID，查找其对应的node信息
如果对应的node节点下线，则为device分配新node*/
func FindNodeNow(dID int) (node *db.Node) {
	node = new(db.Node)
	nodeid, err := FindNodeIDRedis(dID)
	if err == nil {
		node = FindNodeInfoRedis(nodeid)
		if node != nil {
			//成功找到节点信息，返回
			return
		}
	}
	nodeid, err = FindNowNodeIDMysql(dID)
	if err == nil {
		node = FindNodeInfoRedis(nodeid)
		//MySQL中对应的node节点在线，此映射关系依旧可用
		if node != nil {
			InsertDNodeRedis(dID, nodeid)
			return node
		}
	}
	//寻找节点失败，开始分配新节点
	newNodeID, err := AllocateNewNode(dID)
	if err != nil {
		log.Log.Warnln("device -FindNodeNow 分配新节点失败 ", err.Error())
		return nil
	}
	node = FindNodeInfoRedis(newNodeID)
	return

}

//在redis中查找nodeID对应的node信息
func FindNodeInfoRedis(nodeID int) (node *db.Node) {
	key := "nodeinfo" + strconv.Itoa(nodeID)
	value, err := db.Rdb.HMGet(db.RedisCtx, key, "addr", "port").Result()
	if err != nil {
		log.Log.Warnln("device -FindNodeRedis 查找node信息失败 ", err.Error())
		return nil
	}
	if value[0] == nil {
		log.Log.Infoln("device -FindNodeRedis Redis中没有此节点 ", nodeID)
		return nil
	}
	node = new(db.Node)
	node.ID = nodeID
	node.Addr = value[0].(string)
	port, _ := strconv.Atoi(value[1].(string))
	node.Port = port
	return
}

//在redis中查找dID对应的nodeid
func FindNodeIDRedis(dID int) (int, error) {
	key := "device" + strconv.Itoa(dID)
	value, err := db.Rdb.Get(db.RedisCtx, key).Result()
	if err != nil {
		log.Log.Infoln("device -FindNodeIdRedis 查找nodeID失败 ", err.Error())
		return 0, err
	}
	id, _ := strconv.Atoi(value)
	return id, nil
}

/* 通过dID，在MySQL中查找nodeID。
注意！此函数适合查找dID所对应当前最新的nodeID，不适合查找过去的node列表 */
func FindNowNodeIDMysql(dID int) (int, error) {
	var device db.Device
	device.ID = dID
	err := db.Mdb.Order("itime desc").First(&device).Error
	if err != nil {
		log.Log.Infoln("device FindNowNodeIDMysql 查找失败 ", err.Error())
		return 0, err
	}
	return device.Nodeid, nil
}

//为device分配新节点，返回新节点的ID
func AllocateNewNode(dID int) (int, error) {
	nodeID, err := Balance()
	if err != nil {
		return 0, err
	}
	var device db.Device
	device.ID = dID
	device.Nodeid = nodeID
	err = InsertDNodeMysql(dID, &device)
	if err != nil {
		return 0, err
	}
	err = InsertDNodeRedis(dID, device.Nodeid)
	if err != nil {
		return 0, err
	}
	return device.Nodeid, nil
}

//负载均衡算法-待优化
func Balance() (int, error) {
	nodes, err := db.Rdb.Keys(db.RedisCtx, "nodeinfo*").Result()
	if err != nil {
		log.Log.Warnln("device -Balance 负载均衡获取nodeinfo出错 ", err.Error())
		return 0, err
	}
	if len(nodes) == 0 {
		log.Log.Infoln("device -Balance Redis中无node节点信息")
		return 0, err
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	num := rand.Intn(len(nodes))
	nodeID := nodes[num][8:]
	return strconv.Atoi(nodeID)
}

//将dID与nodeID的映射写入mysql
func InsertDNodeMysql(dID int, device *db.Device) error {
	err := db.Mdb.Create(device).Error
	if err != nil {
		log.Log.Warnln("device -InsertDNodeMysql 写入失败", err.Error())
		return err
	}
	return nil
}

//将dID与nodeID的映射写入Redis，传入dID和nodeID
func InsertDNodeRedis(dID int, nodeID int) error {
	key := "device" + strconv.Itoa(dID)

	_, err := db.Rdb.Set(db.RedisCtx, key, nodeID, 0).Result()
	if err != nil {
		log.Log.Warnln("device -InsertDNodeRedis 插入失败 ", err.Error())
		return err
	}
	return nil
}
