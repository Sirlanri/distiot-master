package db

import (
	"github.com/Sirlanri/distiot-master/server/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	//全局DB指针
	db *sqlx.DB
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
