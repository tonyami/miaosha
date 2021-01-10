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
	goodsList, err := goodsService.GetList(page)
	json2(c, goodsList, err)
}

func GetGoods(c *gin.Context) {
	goodsId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		json2(c, nil, code.GoodsNotFound)
		return
	}
	goods, err := goodsService.Get(goodsId)
	if err != nil {
		json2(c, nil, err)
		return
	}
	json2(c, goods.ToVO(), nil)
}
