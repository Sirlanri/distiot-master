package db

import (
	"time"

	"github.com/Sirlanri/distiot-master/server/config"
	"github.com/Sirlanri/distiot-master/server/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	//MySQL数据库指针
	Mdb *gorm.DB
)

func init() {
	connectMysqlByGorm()
	connectRedis()
}

func connectMysqlByGorm() {
	var err error
	dsn := config.Config.MysqlUrl
	Mdb, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Log.Errorln("server-db MySQL连接失败", err.Error())
		return
	}
	err = Mdb.Error
	if err != nil {
		log.Log.Errorln("server-db MySQL ping失败", err.Error())
		return
	}
	log.Log.Infoln("server-db MySQL连接成功")

}

//MySQL内的数据模型
//Node节点表
type Node struct {
	ID   int    `gorm:"primary_key"`
	Addr string `grom:"type:varchar(511)"`
	Port int    `grom:"type:int(0)"`
}

//设备表
type Device struct {
	ID     int       `gorm:"int(0)"`
	Nodeid int       `gorm:"int(0)"`
	Itime  time.Time `gorm:"autoCreateTime:milli"`
}
