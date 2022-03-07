package mqtt

import (
	"encoding/json"

	"github.com/Sirlanri/distiot-master/server/log"
)

//SendMqtt 通过mqtt发送JSON。topic首位不需要添加 '/'
func SendMqttJson(topic string, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Log.Warnln("server-mqtt 发送json 格式化出错", err.Error())
		return
	}
	go func() {
		token := MqClient.Publish(Topic+topic, 2, false, data)
		err = token.Error()
		if err != nil {
			log.Log.Warnln("server-mqtt 发送json 失败", err.Error())
			return
		}
		token.Wait()
	}()

}

//SendMqttString 通过mqtt发送消息，文本格式
func SendMqttString(topic, payload string) {
	go func() {
		token := MqClient.Publish(Topic+topic, 2, false, payload)
		err := token.Error()
		if err != nil {
			log.Log.Warnln("server-mqtt 发送文本出错", err.Error())
			return
		}
		token.Wait()
	}()
}

//SendMqttInfo 通过mqtt发送info调试消息，文本格式
func SendMqttInfo(payload string) {
	go func() {
		token := MqClient.Publish(Topic+"log/master/info", 2, false, payload)
		err := token.Error()
		if err != nil {
			log.Log.Warnln("server-mqtt 发送info失败", err.Error())
			return
		}
		token.Wait()
	}()
}

//SendMqttWarn 通过mqtt发送warn信息，文本格式
func SendMqttWarn(payload string) {
	go func() {
		token := MqClient.Publish(Topic+"log/master/warn", 2, false, payload)
		err := token.Error()
		if err != nil {
			log.Log.Warnln("server-mqtt 发送info失败", err.Error())
			return
		}
		token.Wait()
	}()
}

//SendMqttError 通过mqtt发送error信息，文本格式
func SendMqttError(payload string) {
	go func() {
		token := MqClient.Publish(Topic+"log/master/error", 2, false, payload)
		err := token.Error()
		if err != nil {
			log.Log.Warnln("server-mqtt 发送info失败", err.Error())
			return
		}
		token.Wait()
	}()
}
