package db

import (
	"github.com/Sirlanri/distiot-master/server/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var (
	//全局DB指针
	db *sqlx.DB
	//MySQL数据库指针
	Mdb *gorm.DB
)

func init() {
	connectMysql()
	connectRedis()
}

//初始化MySQL数据库连接
func connectMysql() {
	var err error
	db, err = sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/distiot-master")
	if err != nil {
		log.Log.Errorln("server- mysql open数据库失败")
		return
	}
	err = db.Ping()
	if err != nil {
		log.Log.Errorln("server- MySQL ping数据库失败")
		return
	}
	log.Log.Infoln("server- MySQL链接完成")

}

func connectMysqlByGorm() {
	var err error
	Mdb, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/distiot")
	if err != nil {
		log.Log.Errorln("server-db MySQL连接失败")
		return
	}
	err = Mdb.DB().Ping()
	if err != nil {
		log.Log.Errorln("server-db MySQL ping失败")
		return
	}
	log.Log.Infoln("server-db MySQL连接成功")

}
