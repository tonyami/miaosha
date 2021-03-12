package main

import (
	"log"
	"miaosha/conf"
	"miaosha/infra/db"
	"miaosha/infra/rdb"
	"miaosha/jobs"
	"miaosha/router"
)

func main() {
	if err := conf.Init("./conf.ini"); err != nil {
		log.Fatalf("conf init failed, err:%v", err)
	}
	if err := db.Init(); err != nil {
		log.Fatalf("db init failed, err:%v", err)
	}
	if err := rdb.Init(); err != nil {
		log.Fatalf("rdb init failed, err:%v", err)
	}
	jobs.Init()
	if err := router.Init(); err != nil {
		log.Fatalf("router init failed, err:%v", err)
	}
}
