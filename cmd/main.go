package main

import (
	"miaosha/conf"
	"miaosha/http"
)

func main() {
	conf.Init()
	http.Init()
}
