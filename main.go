package main

import (
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	log.Log.Infoln("master节点 启动完成")
}
