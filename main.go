package main

import (
	"time"

	httphandler "github.com/Sirlanri/distiot-master/httpHandler"
	_ "github.com/Sirlanri/distiot-master/node"
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/device"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	start := time.Now()
	go httphandler.IrisInit()
	elapsed := time.Since(start)
	log.Log.Debugln("执行时间 ", elapsed)

	log.Log.Infoln("master节点 启动完成")
	for i := 0; i < 20; i++ {
		res, err := device.Balance()
		if err != nil {
			log.Log.Errorln("节点负载均衡失败", err.Error())
		}
		log.Log.Debugln(res)
	}
	select {}
}
