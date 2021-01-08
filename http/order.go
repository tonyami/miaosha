package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/model"
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
	if err != nil {
		JSON2(c, nil, err)
		return
	}
	JSON2(c, orderId, nil)
}
