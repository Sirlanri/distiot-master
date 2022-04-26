package db

import (
	"context"

	"github.com/Sirlanri/distiot-master/server/config"
	"github.com/Sirlanri/distiot-master/server/log"
	"github.com/go-redis/redis/v8"
)

//Redis服务器上下文ctx
var RedisCtx = context.Background()

//Redis服务器客户端
var Rdb *redis.Client

func connectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddr,
		Username: config.Config.RedisName,
		Password: config.Config.RedisPW,
		DB:       config.Config.RedisDB,
	})

	str, err := Rdb.Ping(RedisCtx).Result()
	if err != nil {
		log.Log.Errorln("server- redis ping失败")
		return
	}
	log.Log.Infoln("server- redis 连接成功", str)

}
