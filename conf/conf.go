package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

var (
	conf   *Config
	Server *ServerConf
	Mysql  *MysqlConf
)

type Config struct {
	Server *ServerConf `json:"server"`
	Mysql  *MysqlConf  `json:"mysql"`
}

type ServerConf struct {
	Port int `json:"port"`
}

type MysqlConf struct {
	DSN      string        `json:"dsn"`
	Idles    int           `json:"idles"`
	Opens    int           `json:"opens"`
	Lifetime time.Duration `json:"lifetime"`
}

func Init(env string) error {
	if env == "test" {
		return load(os.Getenv("MIAOSHA_ROOT") + "/conf.json")
	}
	return load("conf.json")
}

func load(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &conf); err != nil {
		return
	}
	Server = conf.Server
	Mysql = conf.Mysql
	return
}
