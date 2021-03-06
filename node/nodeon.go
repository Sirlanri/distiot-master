//node 已知的节点上线
package node

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Sirlanri/distiot-master/server/db"
	"github.com/Sirlanri/distiot-master/server/log"
)

//根据node的id，在redis和mysql中查找，若只在MySQL存在，则存入Redis （此方法不合逻辑，暂时废弃）
func findNodeByid(id int) (addr, port string, err error) {
	addr, port, err = FindNodeRds(id)
	if err != nil {
		node, err := FindNodeMysql(id)
		//Redis无数据，MySQL有数据，则将数据存入Redis
		if err == nil {
			addr = node.Addr
			port = strconv.Itoa(node.Port)
			InsertNodeRedis(&node)
			return addr, port, nil
		}
		err = errors.New("node- FindNodeByid redis与mysql中均无数据 " + strconv.Itoa(id))
	}
	return
}

//为新node分配id，将id发送至node节点
func DistributeID(id int, addr string, port int) {
	log.Log.Infoln("分配ID ", id, addr, port)
	params := url.Values{}
	Url, err := url.Parse("http://" + addr + ":" + strconv.Itoa(port) + "/node/allocateid")
	if err != nil {
		log.Log.Warnln("node- DistributeID url解析失败 ", err.Error())
		return
	}
	params.Set("id", strconv.Itoa(id))
	Url.RawQuery = params.Encode()
	res, err := http.Get(Url.String())
	if err != nil {
		log.Log.Warnln("node- DistributeID http get失败 ", err.Error())
		return
	}
	log.Log.Infoln("node- DistributeID 返回结果：", res.StatusCode)
}

//确认节点上线 发送至node
func NodeOnConfirm(node *db.Node) {
	log.Log.Debugln("node上线确认：", strconv.Itoa(node.ID))
	Url, _ := url.Parse("http://" + node.Addr + ":" + strconv.Itoa(node.Port) + "/node/onlineconfirm")
	res, err := http.Get(Url.String())
	if err != nil {
		log.Log.Warnln("node- NodeOnConfirm 确认节点上线失败 ", err.Error())
		return
	}
	log.Log.Infoln("node- NodeOnConfirm 返回结果：", res.StatusCode)
}

//主动检测node是否还活着。如果err不为nil，则检测失败
func NodeOnCheck(node *db.Node) (bool, error) {
	log.Log.Debugln("正在确定节点存活 ", strconv.Itoa(node.ID))
	Url, _ := url.Parse(node.Addr + ":" + strconv.Itoa(node.Port) + "/node/ruok")
	res, err := http.Get(Url.String())
	if err != nil {
		log.Log.Warnln("node- NodeOnCheck 确认节点存活失败 ", err.Error())
		return false, err
	}
	defer res.Body.Close()
	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Log.Warnln("node- NodeOnCheck IO读取失败 ", err.Error())
		return false, err
	}
	log.Log.Infoln("node- NodeOnConfirm 返回结果：", res.StatusCode)
	if string(s) == "ok" {
		log.Log.Infoln("node- NodeOnCheck 节点存活 ", strconv.Itoa(node.ID))
		return true, nil
	} else {
		log.Log.Infoln("node- NodeOnCheck 节点不存活 ", strconv.Itoa(node.ID))
		return false, nil
	}
}
