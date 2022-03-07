package db

import (
	"context"

	"github.com/Sirlanri/distiot-master/server/log"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func connectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Log.Errorln("server- redis ping失败")
		return
	}
	log.Log.Infoln("server- redis 连接成功")

}
