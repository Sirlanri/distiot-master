package node

import (
	"errors"
	"strconv"
	"time"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

//根据node的id，在Redis获取该node的地址、端口信息
func FindNodeRds(id int) (addr string, port string, err error) {
	ids := "nodeinfo" + strconv.Itoa(id)
	cmd := db.Rdb.HMGet(db.RedisCtx, ids, addr, port)
	value, err := cmd.Result()
	if err != nil || value[0] == nil {
		//log.Log.Warnln("node- FindNodeRds 无法获取该node信息:", id)
		err = errors.New("node- FindNodeRds 无法获取该node信息: " + ids)
		return "", "0", err
	}
	var arr [2]string
	for index, v := range value {
		arr[index] = v.(string)
	}
	return arr[0], arr[1], nil
}

//根据node的id，在MySQL中获取该node的地址、端口信息
func FindNodeMysql(id int) (node db.Node, err error) {
	db.Mdb.First(&node, id)
	if node.Addr == "" {
		log.Log.Warnln("node MySQL查找该节点信息失败:", id)
		return node, errors.New("node MySQL查找该节点信息失败")
	}
	log.Log.Debugln("MySQL查找成功 ", node)
	return node, nil
}

//将nodes的单个数据写入redis
func InsertNodeRedis(node *db.Node) error {
	key := "nodeinfo" + strconv.Itoa(node.ID)
	_, err := db.Rdb.HMSet(db.RedisCtx, key, "addr", node.Addr,
		"port", node.Port, "id", node.ID).Result()
	if err != nil {
		log.Log.Warnln("node- InsertNodeRedis node节点信息写入Redis失败 ", err.Error())
		return errors.New("node- InsertNodeRedis node节点信息写入Redis失败 ")
	}
	err = db.Rdb.Expire(db.RedisCtx, key, time.Second*2).Err()
	if err != nil {
		log.Log.Warnln("node- InsertNodeRedis node节点信息写入Redis设置过期时间失败 ", err.Error())
		return errors.New("node- InsertNodeRedis node节点信息写入Redis设置过期时间失败 ")
	}
	return nil
}

func InsertHeartBeatRedis(hbData *HeartBeat) error {
	key := "nodeinfo" + strconv.Itoa(hbData.ID)
	_, err := db.Rdb.HMSet(db.RedisCtx, key, "cpu", hbData.CPU,
		"mem", hbData.Mem, "disk", hbData.Disk).Result()
	if err != nil {
		log.Log.Warnln("node- InsertHeartBeatRedis node心跳信息写入Redis失败 ", err.Error())
		return errors.New("node- InsertHeartBeatRedis node心跳信息写入Redis失败 ")
	}
	err = db.Rdb.Expire(db.RedisCtx, key, time.Second*2).Err()
	if err != nil {
		log.Log.Warnln("node- InsertHeartBeatRedis node心跳信息写入Redis设置过期时间失败 ", err.Error())
		return errors.New("node- InsertHeartBeatRedis node心跳信息写入Redis设置过期时间失败 ")
	}
	return nil
}

//将node的单个数据写入MySQL，返回此node的id （未来优化性能入手点之一）
func InsertNodeMysql(addr string, port int) (int, error) {
	node := &db.Node{
		ID:   0,
		Addr: addr,
		Port: port,
	}
	res := db.Mdb.Create(node)
	if res.Error != nil {
		log.Log.Warnln("node- InsertNodeMysql 插入数据失败 ", res.Error)
		return 0, res.Error
	}
	log.Log.Debugln("写入成功", res.RowsAffected)
	return node.ID, nil
}

//更新已有node的信息至MySQL
func UpdateNodeMysql(node *db.Node) error {
	res := db.Mdb.Save(&node)
	if res.Error != nil {
		log.Log.Warnln("node- UpdateNodeMysql 更新数据失败 ", res.Error)
		return res.Error
	}
	return nil
}
