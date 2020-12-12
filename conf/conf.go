package conf

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

var Conf *Config

type Config struct {
	Server Server `json:"server"`
	Mysql  *Mysql `json:"mysql"`
}

type Server struct {
	Port int `json:"port"`
}

type Mysql struct {
	DSN      string        `json:"dsn"`
	Idles    int           `json:"idles"`
	Opens    int           `json:"opens"`
	Lifetime time.Duration `json:"lifetime"`
}

func Init(path string) error {
	if path == "" {
		return load("conf.json")
	}
	if path == "test" {
		return load("D:\\dev\\code\\miaosha\\cmd\\conf.json")
	}
	return load(path)
}

func load(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, &Conf); err != nil {
		return
	}
	return
}
