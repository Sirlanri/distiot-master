package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

//用于实现订阅的接口
type Handler interface {
	deal(client mqtt.Client, msg mqtt.Message)
}

func SubFunc(subTopic string, handler Handler) {
	MqClient.Subscribe(subTopic, 2, msgHandler)
}
