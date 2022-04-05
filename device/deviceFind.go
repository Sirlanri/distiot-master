package device

import (
	"time"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
	"gorm.io/gorm"
)

/* 在MySQL中 通过的ID寻找对应的node信息及开始保存在此node的时间戳，
nodes信息为之前绑定的节点，返回整个列表*/
func FindNodeIDMysql(dID int) (nodes []NodeWithTime) {
	var devices []db.Device
	//使用事务查询
	session := gorm.Session{
		SkipDefaultTransaction: true,
	}
	tx := db.Mdb.Session(&session)
	err := tx.Find(&devices, dID).Error
	if err != nil {
		log.Log.Warnln("device -FindNodeIDMysql Find dID失败", err)
		tx.Rollback()
		return nil
	}
	for _, d := range devices {
		var node db.Node
		var nodewithtime NodeWithTime
		err = tx.First(&node, d.Nodeid).Error
		if err != nil {
			log.Log.Warnln("device -FindNodeIDMysql Find node信息失败", err)
			tx.Rollback()
			return nil
		}
		nodewithtime.Addr = node.Addr
		nodewithtime.Port = node.Port
		nodewithtime.Itime = d.Itime
		nodes = append(nodes, nodewithtime)
	}
	return
}

type NodeWithTime struct {
	Itime time.Time
	Addr  string
	Port  int
}
