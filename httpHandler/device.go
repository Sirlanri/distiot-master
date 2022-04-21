package httphandler

import (
	"github.com/Sirlanri/distiot-master/server/device"
	"github.com/Sirlanri/distiot-master/server/log"
	"github.com/kataras/iris/v12"
)

func GetNodeHandler(con iris.Context) {
	id, err := con.URLParamInt("id")
	if err != nil {
		log.Log.Warnln("GetNodeHandler 获取节点id失败", err.Error())
		con.StatusCode(401)
		con.WriteString("传入格式错误，没有id")
		return
	}
	node := device.FindNodeNow(id)
	if node == nil {
		con.StatusCode(500)
		con.WriteString("服务器错误")
		return
	}
	_, err = con.JSON(node)
	if err != nil {
		log.Log.Warnln("GetNodeHandler json格式化失败", err.Error())
		con.StatusCode(500)
		con.WriteString("服务器错误")
		return
	}
}
