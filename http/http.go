package http

import (
	"flag"
	"github.com/gin-gonic/gin"
	"miaosha/http/middleware"
	"miaosha/service/goods"
	"miaosha/service/user"
)

var (
	p        string
	router   *gin.Engine
	userSrv  *user.Service
	goodsSrv *goods.Service
)

func initService() {
	userSrv = user.New()
	goodsSrv = goods.New()
}

func initRouter() {
	router.POST("/user/register", userRegister)
	router.POST("/user/login", userLogin)
	router.Use(middleware.Auth(userSrv))
	router.GET("/user/info", userInfo)
	router.GET("/goods", getGoodsList)
	router.GET("/goods/:goodsId", getGoodsDetail)
}

func Init() error {
	initService()
	router = gin.Default()
	initRouter()
	return router.Run(":" + p)
}

func init() {
	flag.StringVar(&p, "p", "", "server port, default: 8080")
	flag.Parse()
	if p == "" {
		p = "8080"
	}
}

type Resp struct {
	Msg  interface{} `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
