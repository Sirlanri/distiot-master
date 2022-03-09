//用于存储节点状态信息
package node

import "github.com/Sirlanri/distiot-master/server/db"

//存储Nodes节点信息的全局变量
var Nodes []db.Node = make([]db.Node, 0)
