//node 已知的节点上线
package node

import (
	"errors"
	"strconv"

	"github.com/Sirlanri/distiot-master/server/log"
)

//根据node的id，在redis和mysql中查找，若只在MySQL存在，则存入Redis （此方法不合逻辑，暂时废弃）
func findNodeByid(id int) (addr, port string, err error) {
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
		err = errors.New("node- FindNodeByid redis与mysql中均无数据 " + strconv.Itoa(id))
	}
	return
}

//为新node分配id，将id发送至node节点
func DistributeID(id int, addr string, port int) {
	log.Log.Infoln("分配ID ", id, addr, port)
}

//确认节点上线 发送至node
func NodeOnConfirm(id int) {
	log.Log.Debugln("node上线确认：", strconv.Itoa(id))
}
