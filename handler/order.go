package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/model"
	"miaosha/mq"
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
	// 预减库存
	stock, err := repository.DecrStock(r.GoodsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	if stock < 0 {
		if err = repository.IncrStock(r.GoodsId); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": SystemErr,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": "商品已售罄",
		})
		return
	}
	// 校验重复秒杀
	uid, _ := c.Get("uid")
	orderId, err := repository.GetOrderIdByUidAndGid(uid.(int64), r.GoodsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	if len(orderId) > 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"msg": "请勿重复秒杀",
		})
		return
	}
	// 异步下单
	if err = mq.OrderPrecreate.Send(uid.(int64), r.GoodsId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func GetOrderResult(c *gin.Context) {
	goodsId, err := strconv.ParseInt(c.Query("goodsId"), 10, 64)
	if err != nil || goodsId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	uid, _ := c.Get("uid")
	orderId, err := repository.GetOrderIdByUidAndGid(uid.(int64), goodsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "秒杀失败",
		})
		return
	}
	if len(orderId) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 0, // 排队中
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  1, // 秒杀成功
			"orderId": orderId,
		})
	}
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
	if err = repository.CloseOrder(order); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "订单取消失败",
		})
		return
	}
	mq.OrderTimeout.Remove(order.OrderId)
	c.JSON(http.StatusOK, nil)
}
