//node 已知的节点上线
package node

import (
	"errors"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

//根据node的id，获取该node的地址，端口信息
func FindNodeid(id string) (addr string, port string, err error) {
	cmd := db.Rdb.LRange(db.RedisCtx, id, 0, 1)
	value, err := cmd.Result()
	if err != nil {
		log.Log.Warnln("node- FindNodeid 无法获取该node信息:", id)
		err = errors.New("node- FindNodeid 无法获取该node信息: " + id)
		return "", "0", err
	}
	return value[0], value[1], nil
}
