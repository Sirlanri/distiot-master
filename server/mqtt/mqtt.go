package mqtt

import (
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

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("distiot_master_" + Createid())

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	MqClient = mqtt.NewClient(opts)
	if token := MqClient.Connect(); token.Wait() && token.Error() != nil {
		log.Log.Errorln("server-mqtt 连接服务器失败", token.Error())
	}
	//sub()

}

//sub 订阅某个主题的消息
func sub() {
	topic := "distiot/online/newnode/"
	token := MqClient.Subscribe(topic, 2, msgHandler)
	token.Wait()
}
