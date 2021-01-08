package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"miaosha/internal/code"
	"time"
)

var redisCli *redis.Client

type Cache struct{}

func Init() {
	var err error
	if redisCli, err = open(conf.Conf.Redis); err != nil {
		panic(err)
	}
}

func Conn() *Cache {
	return &Cache{}
}

func (*Cache) Get(key string) (value string, err error) {
	var cmd *redis.StringCmd
	if cmd = redisCli.Get(context.Background(), key); cmd.Err() != nil && cmd.Err() != redis.Nil {
		log.Printf("【Redis】Get Failed: %s", cmd)
		err = code.SystemErr
		return
	}
	value = cmd.Val()
	return
}

func (*Cache) Set(key string, value interface{}, expire int) (err error) {
	var cmd *redis.BoolCmd
	if cmd = redisCli.SetNX(context.Background(), key, value, time.Duration(expire)*time.Second); cmd.Err() != nil && cmd.Err() != redis.Nil {
		log.Printf("【Redis】Set Failed: %s", cmd)
		err = code.SystemErr
		return
	}
	return
}

func (*Cache) Expire(key string, expire int) (err error) {
	var cmd *redis.BoolCmd
	if cmd = redisCli.Expire(context.Background(), key, time.Duration(expire)*time.Second); cmd.Err() != nil && cmd.Err() != redis.Nil {
		log.Printf("【Redis】Expire Failed: %s", cmd)
		err = code.SystemErr
		return
	}
	return
}

func open(c *conf.Redis) (cli *redis.Client, err error) {
	cli = redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       0,
	})
	err = cli.Ping(context.Background()).Err()
	return
}
