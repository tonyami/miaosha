package main

import (
	"miaosha/conf"
	"miaosha/http"
)

func main() {
	if err := conf.Init(""); err != nil {
		panic(err)
	}
	http.Init(conf.Conf)
	if err := http.Run(conf.Conf); err != nil {
		panic(err)
	}
}
