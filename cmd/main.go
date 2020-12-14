package main

import (
	"miaosha/http"
)

func main() {
	if err := http.Init(); err != nil {
		panic(err)
	}
}
