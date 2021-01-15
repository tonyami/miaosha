package main

import (
	"miaosha/conf"
	"miaosha/http"
	"miaosha/internal/cache"
	"miaosha/internal/db"
	"miaosha/scheduler"
)

func main() {
	conf.Init()
	db.Init()
	cache.Init()
	scheduler.Init()
	http.Init()
}
