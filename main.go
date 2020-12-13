package main

import (
	"miaosha/api"
	"miaosha/conf"
)

func main() {
	if err := conf.Init(""); err != nil {
		panic(err)
	}
	api.Init()
	if err := api.Run(); err != nil {
		panic(err)
	}
}
