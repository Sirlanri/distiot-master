package mqtt

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sirlanri/distiot-master/server/log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("主题: %s\n", msg.Topic())
	fmt.Printf("消息内容: %s\n", msg.Payload())

}

var msgHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Log.Debugln("接收到指令：", string(msg.Payload()))
}

//全局MQ客户端 通过此发送接收消息
var MqClient mqtt.Client

//全局topic distiot/
const Topic = "distiot/"

//Createid 生成唯一名称
func Createid() string {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	id := u1.String()
	return id[:9]
}

//初始化
func init() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.ri-co.cn:1883").SetClientID("distiot_master_" + Createid())

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	MqClient = mqtt.NewClient(opts)
	if token := MqClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub()

}

//sub 订阅某个主题的消息
func sub() {
	topic := "iot1/server/res"
	token := MqClient.Subscribe(topic, 2, msgHandler)
	token.Wait()
}

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
