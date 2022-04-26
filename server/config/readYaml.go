package config

import (
	"io/ioutil"

	"github.com/Sirlanri/distiot-master/server/log"
	"gopkg.in/yaml.v3"
)

func init() {
	ReadYaml()
}

type Conf struct {
	HttpPort  string `yaml:"httpport"`
	RedisAddr string `yaml:"redisAddr"`
	RedisName string `yaml:"redisName"`
	RedisPW   string `yaml:"redisPW"`
	RedisDB   int    `yaml:"redisDB"`
	MysqlUrl  string `yaml:"mysqlUrl"`
}

//全局配置文件
var Config Conf

func ReadYaml() {
	buf, err := ioutil.ReadFile("masterconf.yaml")
	if err != nil {
		log.Log.Warnln("server-config ReadYaml 读取配置文件失败", err.Error())
		return
	}
	err = yaml.Unmarshal(buf, &Config)
	if err != nil {
		log.Log.Errorln("配置文件读取失败", err.Error())
	}

}
