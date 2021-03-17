package router

import (
	"github.com/gin-gonic/gin"
	"miaosha/handler"
	"miaosha/middleware"
)

func Init() (err error) {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimit())
	router.GET("/goods/stock/init", handler.InitGoodsStock)
	router.GET("/code/sms", handler.GetSmsCode)
	router.POST("/user/login", handler.UserLogin)
	router.GET("/goods/list", handler.GetGoodsList)
	router.GET("/goods", handler.GetGoods)
	router.Use(middleware.Auth())
	router.GET("/user/info", handler.GetUserInfo)
	router.GET("/order", handler.GetOrder)
	router.GET("/order/list", handler.GetOrderList)
	router.POST("/order", handler.CreateOrder)
	router.GET("/order/result", handler.GetOrderResult)
	router.POST("/order/cancel", handler.CancelOrder)
	return router.Run()
}
