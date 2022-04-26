package config

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Sirlanri/distiot-master/server/log"
	"gopkg.in/yaml.v3"
)

func init() {
	ReadYaml()
}

type Conf struct {
	HttpPort string `yaml:"httpport"`
}

//全局配置文件
var Config Conf

func ReadYaml() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index+1]
	fullPath := path + "masterconf.yaml"
	buf, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Log.Warnln("server-config ReadYaml 读取配置文件失败", err.Error())
		return
	}
	err = yaml.Unmarshal(buf, &Config)
	if err != nil {
		log.Log.Errorln("配置文件读取失败", err.Error())
	}

}
