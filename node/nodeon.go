//node 已知的节点上线
package node

import (
	"errors"
	"strconv"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

//根据node的id，依次在redis和mysql中查找，若只在MySQL存在，则存入Redis
func FindNodeByid(id string) (addr, port string, err error) {
	addr, port, err = FindNodeRds(id)
	if err != nil {
		node, err := FindNodeMysql(id)
		//Redis无数据，MySQL有数据，则将数据存入Redis
		if err == nil {
			addr = node.Addr
			port = strconv.Itoa(node.Port)
			InsertNodeRedis(&node)
			return addr, port, nil
		}
		err = errors.New("node- FindNodeByid redis与mysql中均无数据 " + id)
	}
	return
}

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
