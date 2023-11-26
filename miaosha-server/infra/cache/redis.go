package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"time"
)

var Client *redis.Client

func Init() {
	c := conf.Conf.Redis
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	Client = redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       0,
	})
	if err := Client.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}
}
