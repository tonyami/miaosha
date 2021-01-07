package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/internal/code"
	"strconv"
)

func GetGoodsList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	goodsList, err := goodsService.GetGoodsList(page)
	if err != nil {
		JSON2(c, nil, err)
		return
	}
	JSON2(c, goodsList, nil)
}

func GetGoods(c *gin.Context) {
	goodsId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		JSON2(c, nil, code.GoodsNotFound)
		return
	}
	goods, err := goodsService.GetGoods(goodsId)
	if err != nil {
		JSON2(c, nil, err)
		return
	}
	JSON2(c, goods.ToVO(), nil)
}
