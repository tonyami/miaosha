package main

import (
	"miaosha/conf"
	"miaosha/http"
)

func main() {
	conf.Init()
	if err := http.Init(); err != nil {
		panic(err)
	}
}
