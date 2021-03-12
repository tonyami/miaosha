package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/jobs"
	"miaosha/model"
	"miaosha/repository"
	"net/http"
	"strconv"
)

func GetOrderList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	uid, _ := c.Get("uid")
	status := c.Query("status")
	list, err := repository.GetOrderList(uid.(int64), status, page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	orders := make([]model.OrderVO, 0)
	for _, order := range list {
		orders = append(orders, order.ToVO())
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	orderId := c.Query("orderId")
	if len(orderId) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	order, err := repository.GetOrderByOrderId(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	uid, _ := c.Get("uid")
	if order.UserId != uid.(int64) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg": "订单不存在",
		})
		return
	}
	c.JSON(http.StatusOK, order.ToVO())
}

func CreateOrder(c *gin.Context) {
	r := new(struct {
		GoodsId int64 `form:"goodsId" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	// 1、查询秒杀商品
	goods, err := repository.GetGoods(r.GoodsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	if goods.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg": "商品不存在",
		})
		return
	}
	// 2、校验秒杀开始时间、结束时间、库存
	if err = goods.Check(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 3、校验是否重复秒杀
	uid, _ := c.Get("uid")
	count, err := repository.CountRepeatableOrder(uid.(int64), goods.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": "请勿重复秒杀",
		})
		return
	}
	// 4、减库存、创建订单
	order := model.NewOrder(uid.(int64), goods)
	if err = repository.CreateOrder(order); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "秒杀失败",
		})
		return
	}
	jobs.GetOrderTimeoutJob().Add(order.OrderId)
	c.JSON(http.StatusOK, gin.H{
		"orderId": order.OrderId,
	})
}

func CancelOrder(c *gin.Context) {
	r := new(struct {
		OrderId string `form:"orderId" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	order, err := repository.GetOrderByOrderId(r.OrderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	if order.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg": "订单不存在",
		})
		return
	}
	uid, _ := c.Get("uid")
	if order.UserId != uid.(int64) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"msg": "forbidden",
		})
		return
	}
	if order.Status != model.Unpaid {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": "订单无法取消",
		})
		return
	}
	if err = repository.CloseOrder(order.OrderId, order.GoodsId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "订单取消失败",
		})
		return
	}
	jobs.GetOrderTimeoutJob().Remove(order.OrderId)
	c.JSON(http.StatusOK, nil)
}
