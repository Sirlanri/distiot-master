//node 已知的节点上线
package node

import (
	"errors"
	"strconv"
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
