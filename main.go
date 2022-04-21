package main

import (
	"time"

	httphandler "github.com/Sirlanri/distiot-master/httpHandler"
	_ "github.com/Sirlanri/distiot-master/node"
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	start := time.Now()
	go httphandler.IrisInit()
	elapsed := time.Since(start)
	log.Log.Debugln("执行时间 ", elapsed)

	log.Log.Infoln("master节点 启动完成")
	select {}
}
