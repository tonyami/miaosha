package conf

import (
	"gopkg.in/ini.v1"
	"os"
)

var Conf = new(Config)

type Config struct {
	DB        *DB        `ini:"db"`
	Redis     *Redis     `ini:"redis"`
	Order     *Order     `ini:"order"`
	RateLimit *RateLimit `ini:"rate_limit"`
}

type DB struct {
	Host     string `ini:"host"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Name     string `ini:"name"`
	Idles    int    `ini:"idles"`
	Opens    int    `ini:"opens"`
	LifeTime int    `ini:"lifetime"`
}

type Redis struct {
	Host     string `ini:"host"`
	Password string `ini:"password"`
}

type Order struct {
	Expire int64 `ini:"expire"`
}

type RateLimit struct {
	Time  int64 `ini:"time"`
	Count int64 `ini:"count"`
}

func Init(file string) (err error) {
	if file == "" {
		initEnv()
	} else {
		err = initFile(file)
	}
	return
}

func initEnv() (c *Config) {
	Conf.DB = &DB{
		Host:     os.Getenv("MIAOSHA_DB_HOST"),
		User:     os.Getenv("MIAOSHA_DB_USER"),
		Password: os.Getenv("MIAOSHA_DB_PASSWORD"),
		Name:     os.Getenv("MIAOSHA_DB_NAME"),
		Idles:    2,
		Opens:    5,
		LifeTime: 2,
	}
	Conf.Redis = &Redis{
		Host:     os.Getenv("MIAOSHA_REDIS_HOST"),
		Password: os.Getenv("MIAOSHA_REDIS_PASSWORD"),
	}
	Conf.Order = &Order{
		Expire: 1800,
	}
	return
}

func initFile(file string) error {
	return ini.MapTo(&Conf, file)
}
