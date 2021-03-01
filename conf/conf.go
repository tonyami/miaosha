package conf

import (
	"gopkg.in/ini.v1"
	"os"
)

var Conf = new(Config)

type Config struct {
	DB    *DB    `ini:"db"`
	Redis *Redis `ini:"redis"`
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
	return
}

func initFile(file string) error {
	return ini.MapTo(&Conf, file)
}
