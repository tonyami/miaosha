package main

import (
	"log"
	"miaosha/conf"
	"miaosha/infra/cache"
	"miaosha/infra/db"
	"miaosha/router"
)

func init() {
	conf.Init("./conf.ini")
	db.Init()
	cache.Init()
}

func main() {
	app := router.Init()
	if err := app.Run(); err != nil {
		log.Fatalf("app run failed, err:%v", err)
	}
}
