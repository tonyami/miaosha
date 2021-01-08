package main

import (
	"miaosha/conf"
	"miaosha/http"
	"miaosha/internal/cache"
	"miaosha/internal/db"
)

func main() {
	conf.Init()
	db.Init()
	cache.Init()
	http.Init()
}
