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
		addr, port, err = FindNodeMysql(id)
	}
	return
}

//根据node的id，在Redis获取该node的地址、端口信息
func FindNodeRds(id string) (addr string, port string, err error) {
	cmd := db.Rdb.LRange(db.RedisCtx, id, 0, 1)
	value, err := cmd.Result()
	if err != nil || len(value) == 0 {
		log.Log.Warnln("node- FindNodeid 无法获取该node信息:", id)
		err = errors.New("node- FindNodeid 无法获取该node信息: " + id)
		return "", "0", err
	}
	return value[0], value[1], nil
}

//根据node的id，在MySQL中获取该node的地址、端口信息
func FindNodeMysql(id string) (addr, port string, err error) {
	var node db.Node
	db.Mdb.First(&node, id)
	if node.Addr == "" {
		log.Log.Warnln("node MySQL查找该节点信息失败:", id)
		return "", "", errors.New("node MySQL查找该节点信息失败")
	}
	log.Log.Debugln(node)
	return node.Addr, strconv.Itoa(node.Port), nil
}
