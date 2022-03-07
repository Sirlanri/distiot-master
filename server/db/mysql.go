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

func connect() {
	var err error
	db, err = sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Log.Errorln("mysql 连接数据库失败")
	}
}
