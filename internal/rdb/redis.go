package rdb

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"time"
)

var redisCli *redis.Client

func Init(c *conf.Redis) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	redisCli = redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       0,
	})
	if err = redisCli.Ping(ctx).Err(); err != nil {
		cancel()
		return
	}
	return
}

func Get(key string) (val string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	if val, err = redisCli.Get(ctx, key).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("redisCli.Get(%s) failed, err: %v", key, err)
			cancel()
		}
	}
	return
}

func Set(key string, val interface{}, expire int) (err error) {
	var b bool
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	if b, err = redisCli.SetNX(ctx, key, val, time.Duration(expire)*time.Second).Result(); err != nil {
		log.Printf("redisCli.SetNX(%s, %v, %d) failed, err: %v", key, val, expire, err)
		cancel()
		return
	}
	if !b {
		err = errors.New("redisCli.SetNX() failed")
	}
	return
}

func Expire(key string, expire int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	var b bool
	if b, err = redisCli.Expire(ctx, key, time.Duration(expire)*time.Second).Result(); err != nil {
		log.Printf("redisCli.Expire(%s, %d) failed, err: %v", key, expire, err)
		cancel()
		return
	}
	if !b {
		err = errors.New("redisCli.Expire() failed")
		return
	}
	return
}
