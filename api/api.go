package api

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/service"
	"strconv"
)

var (
	e           *gin.Engine
	userService *service.UserService
)

func initService() {
	userService = service.NewUserService()
}

func Init() {
	e = gin.Default()
	initService()
	initRouter()
}

func Run() (err error) {
	if conf.Server.Port == 0 {
		conf.Server.Port = 8080
	}
	port := strconv.Itoa(conf.Server.Port)
	return e.Run(":" + port)
}
