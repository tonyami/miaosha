package main

import (
	"log"
	"miaosha/apis"
	"miaosha/conf"
	"miaosha/internal/db"
	"miaosha/internal/rdb"
	"miaosha/jobs"
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
	if err := apis.Init(); err != nil {
		log.Fatalf("api init failed, err:%v", err)
	}
}
