package node

import (
	"errors"
	"strconv"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

//根据node的id，在Redis获取该node的地址、端口信息
func FindNodeRds(id string) (addr string, port string, err error) {
	cmd := db.Rdb.LRange(db.RedisCtx, id, 0, 1)
	value, err := cmd.Result()
	if err != nil || len(value) == 0 {
		//log.Log.Warnln("node- FindNodeRds 无法获取该node信息:", id)
		err = errors.New("node- FindNodeRds 无法获取该node信息: " + id)
		return "", "0", err
	}
	return value[0], value[1], nil
}

//根据node的id，在MySQL中获取该node的地址、端口信息
func FindNodeMysql(id string) (node db.Node, err error) {
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
	_, err := db.Rdb.RPush(db.RedisCtx, strconv.Itoa(node.Id),
		node.Addr, node.Port).Result()
	if err != nil {
		log.Log.Warnln("node- InsertNodeRedis node节点信息写入Redis失败 ", err.Error())
		return errors.New("node- InsertNodeRedis node节点信息写入Redis失败 ")
	}
	return nil
}

//将node的单个数据写入MySQL，返回此node的id （未来优化性能入手点之一）
func InsertNodeMysql(addr string, port int) (int, error) {
	node := &db.Node{
		Id:   0,
		Addr: addr,
		Port: port,
	}
	res := db.Mdb.Create(node)
	if res.Error != nil {
		log.Log.Warnln("node- InsertNodeMysql 插入数据失败 ", res.Error)
		return 0, res.Error
	}
	log.Log.Debugln("写入成功", res.RowsAffected)
	return node.Id, nil
}
