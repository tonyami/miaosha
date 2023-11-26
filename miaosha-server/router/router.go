package router

import (
	"github.com/gin-gonic/gin"
	"miaosha/handler"
	"miaosha/middleware"
	"miaosha/mq"
	"miaosha/service/goods"
	"miaosha/service/order"
	"miaosha/service/user"
)

var (
	goodsHandler *handler.GoodsHandler
	orderHandler *handler.OrderHandler
	userHandler  *handler.UserHandler
)

func initService() {
	defer mq.Init()
	goods.InitService()
	order.InitService()
	user.InitService()
}

func initHandler() {
	goodsHandler = handler.InitGoodsHandler()
	orderHandler = handler.InitOrderHandler()
	userHandler = handler.InitUserHandler()
}

func initRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimit())
	router.GET("/code/sms", userHandler.GetSmsCode)
	router.POST("/user/login", userHandler.UserLogin)
	router.GET("/goods/list", goodsHandler.GetGoodsList)
	router.GET("/goods", goodsHandler.GetGoods)
	router.GET("/goods/stock/init", goodsHandler.InitGoodsStock)
	router.Use(middleware.Auth())
	router.GET("/user/info", userHandler.GetUserInfo)
	router.GET("/order", orderHandler.GetOrder)
	router.GET("/order/list", orderHandler.GetOrderList)
	router.POST("/miaosha", orderHandler.Miaosha)
	router.GET("/miaosha/result", orderHandler.GetMiaoshaResult)
	router.POST("/order/close", orderHandler.CloseOrder)
}

func Init() (router *gin.Engine) {
	initService()
	initHandler()
	router = gin.Default()
	initRouter(router)
	mq.Run()
	return
}
