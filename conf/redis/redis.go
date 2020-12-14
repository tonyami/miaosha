package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"miaosha/conf"
	"os"
	"sync"
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
	addr := os.Getenv(conf.ENV_MIAOSHA_REDIS)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	return rdb
}
