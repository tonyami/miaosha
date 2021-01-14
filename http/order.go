package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/model"
	"strconv"
)

func Miaosha(c *gin.Context) {
	r := new(struct {
		GoodsId int64 `form:"goodsId" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	user, ok := c.Get(conf.User)
	if !ok {
		return
	}
	u := user.(*model.User)
	orderId, err := orderService.Miaosha(u.Id, r.GoodsId)
	json2(c, orderId, err)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		return
	}
	user, ok := c.Get(conf.User)
	if !ok {
		return
	}
	u := user.(*model.User)
	order, err := orderService.Get(id, u.Id)
	json2(c, order, err)
}

func GetOrderList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	status := c.Query("status")
	user, ok := c.Get(conf.User)
	if !ok {
		return
	}
	u := user.(*model.User)
	orders, err := orderService.GetList(u.Id, page, status)
	json2(c, orders, err)
}

func OrderCancel(c *gin.Context) {
	r := new(struct {
		Id string `form:"id" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	user, ok := c.Get(conf.User)
	if !ok {
		return
	}
	u := user.(*model.User)
	err := orderService.Cancel(r.Id, u.Id)
	json2(c, nil, err)
}
