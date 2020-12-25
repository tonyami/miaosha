package conf

import (
	"os"
)

const (
	TokenPrefix             = "token:"
	TokenExpire             = 2 * 3600
	UserSession             = "user"
	GoodsSize               = 10
	MiaoshaStatusNotStarted = 0
	MiaoshaStatusOnGoing    = 1
	MiaoshaStatusFinished   = 2
)

var Conf *Config

type Config struct {
	DbAddr        string
	DbUser        string
	DbPassword    string
	RedisAddr     string
	RedisPassword string
}

func Init() {
	initEnv()
}

func initEnv() {
	Conf = &Config{
		DbAddr:        os.Getenv("MIAOSHA_DB_ADDR"),
		DbUser:        os.Getenv("MIAOSHA_DB_USER"),
		DbPassword:    os.Getenv("MIAOSHA_DB_PASSWORD"),
		RedisAddr:     os.Getenv("MIAOSHA_REDIS_ADDR"),
		RedisPassword: os.Getenv("MIAOSHA_REDIS_PASSWORD"),
	}
}
