package apis

import (
	"github.com/gin-gonic/gin"
	_ "miaosha/core/goods"
	_ "miaosha/core/order"
	_ "miaosha/core/user"
	"miaosha/service"
)

var (
	userService  service.UserService
	goodsService service.GoodsService
	orderService service.OrderService
)

func initService() {
	userService = service.GetUserService()
	goodsService = service.GetGoodsService()
	orderService = service.GetOrderService()
}

func initRouter(router *gin.Engine) {
	router.Use(Cors())
	router.GET("/code/sms", SendSmsCode)
	router.POST("/user/login", Login)
	router.GET("/goods/list", GetGoodsList)
	router.GET("/goods", GetGoods)
	router.Use(Auth())
	router.GET("/user", GetUserInfo)
	router.GET("/order", GetOrder)
	router.GET("/order/list", GetOrderList)
	router.POST("/order", CreateOrder)
	router.POST("/order/cancel", CancelOrder)
}

func Init() (err error) {
	initService()
	router := gin.Default()
	initRouter(router)
	return router.Run()
}
