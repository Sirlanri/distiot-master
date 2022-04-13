//node相关的mqtt处理
package node

import (
	"encoding/json"
	"errors"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
	"github.com/Sirlanri/distiot-master/server/mqtt"
	mq "github.com/eclipse/paho.mqtt.golang"
)

func init() {
	ListenNode()
}

//绑定目mqtt监听
func ListenNode() {
	//已知节点上线
	token1 := mqtt.MqClient.Subscribe(mqtt.Topic+"online/oldnode", 2, OldNodeonHandler)
	//新节点上线
	token2 := mqtt.MqClient.Subscribe(mqtt.Topic+"online/newnode", 2, NewNodeonHandler)
	//接收node节点的心跳数据
	token3 := mqtt.MqClient.Subscribe(mqtt.Topic+"heartbeat/node", 2, HeartbeatHandler)
	token1.Wait()
	token2.Wait()
	token3.Wait()
}

func HeartbeatHandler(client mq.Client, msg mq.Message) {
	payload := msg.Payload()
	data := FormatPayloadHeartbeat(&payload)
	if data == nil {
		return
	}
	//更新Redis
	InsertHeartBeatRedis(data)

}

//旧node节点上线
func OldNodeonHandler(client mq.Client, msg mq.Message) {
	payload := msg.Payload()
	data, err := FormatPayloadOldNode(&payload)
	if err != nil {
		return
	}
	node, err := FindNodeMysql(data.ID)
	if err != nil {
		return
	}
	if node.Addr != data.Addr || node.Port != data.Port {
		//节点信息发生变化，将新数据写入MySQL
		UpdateNodeMysql(data)
	}
	//更新Redis上线
	err = InsertNodeRedis(data)

	if err == nil {
		NodeOnConfirm(data)
	}

}

//新node节点上线，分配id
func NewNodeonHandler(client mq.Client, msg mq.Message) {
	log.Log.Debugln("接收到消息")
	payload := msg.Payload()
	data, err := FormatPayloadNewNode(&payload)
	if err != nil {
		return
	}
	id, err := InsertNodeMysql(data.Addr, data.Port)
	if err != nil {
		return
	}
	DistributeID(id, data.Addr, data.Port)
}

//格式化mqtt传入数据-新节点
func FormatPayloadNewNode(payload *[]byte) (*NewNodeOn, error) {
	var nodeStruct NewNodeOn
	err := json.Unmarshal(*payload, &nodeStruct)
	if err != nil {
		log.Log.Warnln("node- FormatPayloadNewNode失败", err.Error())
		return nil, errors.New("node- FormatPayload失败")
	}
	return &nodeStruct, nil
}

//格式化mqtt传入数据-旧节点
func FormatPayloadOldNode(payload *[]byte) (*db.Node, error) {
	var nodeStruct db.Node
	err := json.Unmarshal(*payload, &nodeStruct)
	if err != nil {
		log.Log.Warnln("node- FormatPayloadOldNode失败", err.Error())
		return nil, errors.New("node- FormatPayload失败")
	}
	return &nodeStruct, nil
}

func FormatPayloadHeartbeat(payload *[]byte) (hbData *HeartBeat) {
	hbData = new(HeartBeat)
	err := json.Unmarshal(*payload, hbData)
	if err != nil {
		log.Log.Warnln("node- FormatPayloadHeartbeat失败", err.Error())
		return nil
	}
	return hbData
}

//新节点上线的mqtt数据结构
type NewNodeOn struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

//有ID的旧节点上线
type OldNodeOn struct {
	ID   int    `json:"id"`
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type HeartBeat struct {
	ID   int `json:"id"`
	CPU  int `json:"cpu"`
	Mem  int `json:"mem"`
	Disk int `json:"disk"`
}
