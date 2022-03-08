package main

import (
	"github.com/Sirlanri/distiot-master/node"
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	log.Log.Infoln("master节点 启动完成")
	addr, port, err := node.FindNodeByid("1")
	if err != nil {
		log.Log.Errorln("出错！", err.Error())
		return
	}
	log.Log.Infoln(addr, port)
}
