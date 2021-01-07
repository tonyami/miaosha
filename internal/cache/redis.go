package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"miaosha/conf"
	"sync"
)

var (
	cli  *redis.Client
	once sync.Once
)

func New(c *conf.Redis) *redis.Client {
	once.Do(func() {
		cli = open(c)
	})
	return cli
}

func open(c *conf.Redis) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       0,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return cli
}
