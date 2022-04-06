package main

import (
	"sync"
	"time"

	_ "github.com/Sirlanri/distiot-master/node"
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	var wg2 sync.WaitGroup
	wg2.Add(1)
	start := time.Now()

	elapsed := time.Since(start)
	log.Log.Debugln("执行时间 ", elapsed)

	log.Log.Infoln("master节点 启动完成")
	wg2.Wait()
}
