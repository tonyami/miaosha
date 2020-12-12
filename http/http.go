package http

import (
	"log"
	"miaosha/conf"
	"miaosha/service/user"
	"net/http"
	"strconv"
)

var (
	userSvc *user.Service
)

func Init(c *conf.Config) {
	initService(c)
	initRouter()
}

func Run(c *conf.Config) (err error) {
	if c.Server.Port == 0 {
		c.Server.Port = 8080
	}
	log.Printf("Server started on port: %d\n", c.Server.Port)
	return http.ListenAndServe(":"+strconv.Itoa(c.Server.Port), nil)
}

func initService(c *conf.Config) {
	userSvc = user.New(c)
}

func initRouter() {
	http.HandleFunc("/user/info", userInfo)
}
