package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/internal/code"
	"miaosha/service/goods"
	"miaosha/service/order"
	"miaosha/service/user"
	"net/http"
)

var (
	userService  *user.Service
	goodsService *goods.Service
	orderService *order.Service
)

func initService() {
	userService = user.New()
	goodsService = goods.New()
	orderService = order.New(goodsService)
}

func initRouter(router *gin.Engine) {
	router.GET("/code/sms", SendSmsCode)
	router.POST("/user/login", Login)
	router.Use(Auth())
	router.GET("/user", GetUser)
	router.GET("/goods", GetGoodsList)
	router.GET("/goods/:id", GetGoods)
	router.POST("/order", Miaosha)
}

func Init() {
	initService()
	router := gin.Default()
	initRouter(router)
	if err := router.Run(); err != nil {
		panic(err)
	}
}

func JSON2(c *gin.Context, data interface{}, err error) {
	if err != nil {
		ec := code.String(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": ec.Code(),
			"msg":  ec.Message(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code.Success.Code(),
		"msg":  code.Success.Message(),
		"data": data,
	})

}
