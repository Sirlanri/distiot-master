//node相关的mqtt处理
package node

import (
	"encoding/json"
	"errors"

	"github.com/Sirlanri/distiot-master/server/log"
	"github.com/Sirlanri/distiot-master/server/mqtt"
	mq "github.com/eclipse/paho.mqtt.golang"
)

//绑定目mqtt监听
func ListenNode() {
	//已知节点上线
	//token1 := mqtt.MqClient.Subscribe(mqtt.Topic+"online/node", 2, NodeonHandler)
	//新节点上线
	token2 := mqtt.MqClient.Subscribe(mqtt.Topic+"online/newnode/", 2, NewNodeonHandler)
	//token1.Wait()
	token2.Wait()

}

//node节点上线
func NodeonHandler(client mq.Client, msg mq.Message) {

}

//新node节点上线，分配id
func NewNodeonHandler(client mq.Client, msg mq.Message) {
	log.Log.Debugln("接收到消息")
	payload := msg.Payload()
	data, err := FormatPayload(&payload)
	if err != nil {
		return
	}
	id, err := InsertNodeMysql(data.Addr, data.Port)
	if err != nil {
		return
	}
	DistributeID(id, data.Addr, data.Port)
}

//格式化mqtt传入数据
func FormatPayload(payload *[]byte) (*NewNodeOn, error) {
	var nodeStruc NewNodeOn
	err := json.Unmarshal(*payload, &nodeStruc)
	if err != nil {
		log.Log.Warnln("node- FormatPayload失败", err.Error())
		return nil, errors.New("node- FormatPayload失败")
	}
	return &nodeStruc, nil
}

//新节点上线的mqtt数据结构
type NewNodeOn struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}
