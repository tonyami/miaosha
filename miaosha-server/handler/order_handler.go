package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/service"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderService service.IOrderService
}

func InitOrderHandler() *OrderHandler {
	return &OrderHandler{orderService: service.GetOrderService()}
}

func (h *OrderHandler) GetOrderList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	uid, _ := c.Get("uid")
	status := c.Query("status")
	list, err := h.orderService.GetOrderList(uid.(int64), status, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, list)
	}

}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderId := c.Query("orderId")
	if len(orderId) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	uid, _ := c.Get("uid")
	orderInfo, err := h.orderService.GetOrderInfoVO(orderId, uid.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, orderInfo)
	}
}

func (h *OrderHandler) Miaosha(c *gin.Context) {
	r := new(struct {
		GoodsId int64 `form:"goodsId" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	uid, _ := c.Get("uid")
	if err := h.orderService.Miaosha(uid.(int64), r.GoodsId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func (h *OrderHandler) GetMiaoshaResult(c *gin.Context) {
	goodsId, err := strconv.ParseInt(c.Query("goodsId"), 10, 64)
	if err != nil || goodsId == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	uid, _ := c.Get("uid")
	result, err := h.orderService.GetMiaoshaReuslt(uid.(int64), goodsId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (h *OrderHandler) CloseOrder(c *gin.Context) {
	r := new(struct {
		OrderId string `form:"orderId" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	uid, _ := c.Get("uid")
	if err := h.orderService.CloseOrder(uid.(int64), r.OrderId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
