package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetOrderList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	status := c.Query("status")
	uid, _ := c.Get("uid")
	orders, err := orderService.GetList(uid.(int64), page, status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	uid, _ := c.Get("uid")
	order, err := orderService.Get(id, uid.(int64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	r := new(struct {
		GoodsId int64 `form:"goodsId" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	uid, _ := c.Get("uid")
	orderId, err := orderService.Create(uid.(int64), r.GoodsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orderId": orderId,
	})
}

func CancelOrder(c *gin.Context) {
	r := new(struct {
		Id int64 `form:"id" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	uid, _ := c.Get("uid")
	err := orderService.Cancel(r.Id, uid.(int64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
