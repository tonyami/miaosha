package main

import (
	"log"
	"miaosha/apis"
	"miaosha/conf"
	"miaosha/internal/db"
	"miaosha/internal/rdb"
)

func main() {
	conf.Init()
	if err := db.Init(conf.Conf.DB); err != nil {
		log.Fatalf("db init failed, err:%v", err)
	}
	if err := rdb.Init(conf.Conf.Redis); err != nil {
		log.Fatalf("rdb init failed, err:%v", err)
	}
	if err := apis.Init(); err != nil {
		log.Fatalf("api init failed, err:%v", err)
	}
}
