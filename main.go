package main

import (
	"sync"

	"github.com/Sirlanri/distiot-master/node"
	_ "github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

func main() {
	var wg2 sync.WaitGroup
	wg2.Add(1)
	node.ListenNode()

	log.Log.Infoln("master节点 启动完成")
	wg2.Wait()
}
