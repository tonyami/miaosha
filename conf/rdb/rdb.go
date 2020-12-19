package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"sync"
	"time"
)

var (
	ctx  = context.Background()
	rdb  *redis.Client
	once sync.Once
)

func New() *redis.Client {
	once.Do(func() {
		rdb = connect()
	})
	return rdb
}

func connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:        conf.Conf.RedisAddr,
		Password:    conf.Conf.RedisPassword,
		DB:          0,
		DialTimeout: 1 * time.Second,
		MaxRetries:  2,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Print(err)
	}
	return rdb
}
