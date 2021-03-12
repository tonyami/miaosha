package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"miaosha/conf"
	"time"
)

var redisCli *redis.Client

func Init() (err error) {
	c := conf.Conf.Redis
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	redisCli = redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       0,
	})
	err = redisCli.Ping(ctx).Err()
	return
}

func Conn() *redis.Client {
	return redisCli
}
