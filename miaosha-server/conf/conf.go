package conf

import (
	"log"

	"gopkg.in/ini.v1"
)

var Conf Config

type Config struct {
	Server    Server    `init:"server"`
	DB        db        `ini:"db"`
	Redis     redis     `ini:"redis"`
	Order     order     `ini:"order"`
	RateLimit rateLimit `ini:"rate_limit"`
}

type Server struct {
	Port string `ini:"port"`
}

type db struct {
	Host     string `ini:"host"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Name     string `ini:"name"`
	Idles    int    `ini:"idles"`
	Opens    int    `ini:"opens"`
}

type redis struct {
	Host     string `ini:"host"`
	Password string `ini:"password"`
}

type order struct {
	Expire int64 `ini:"expire"`
}

type rateLimit struct {
	Time  int64  `ini:"time"`
	Count int64  `ini:"count"`
	Port  string `ini:"port"`
}

func Init(file string) {
	if err := ini.MapTo(&Conf, file); err != nil {
		log.Fatal(err)
	}
	log.Printf("config: %+v", Conf)
}
