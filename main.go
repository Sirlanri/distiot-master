package main

import (
	"sync"
	"time"

	"github.com/Sirlanri/distiot-master/device"
	_ "github.com/Sirlanri/distiot-master/node"
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	var wg2 sync.WaitGroup
	wg2.Add(1)
	start := time.Now()
	node := device.FindNodeNow(2)
	log.Log.Debugln("：1 不需要插入", node)
	node = device.FindNodeNow(3)
	log.Log.Debugln("：1 插入redis 3", node)
	node = device.FindNodeNow(1)
	log.Log.Debugln("：2 插入redis 1", node)

	node = device.FindNodeNow(4)
	log.Log.Debugln("随机 1 2", node)

	elapsed := time.Since(start)
	log.Log.Debugln("执行时间 ", elapsed)

	log.Log.Infoln("master节点 启动完成")
	wg2.Wait()
}
