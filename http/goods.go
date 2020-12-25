package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf/errmsg"
	"net/http"
	"strconv"
)

func getGoodsList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			Msg: errmsg.InvalidParameter,
		})
		return
	}
	goodsList, err := goodsSrv.GetGoodsList(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Msg: err,
		})
		return
	}
	c.JSON(http.StatusOK, Resp{
		Data: goodsList,
	})
}

func getGoodsDetail(c *gin.Context) {
	goodsId, err := strconv.ParseInt(c.Param("goodsId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Resp{
			Msg: errmsg.InvalidParameter,
		})
		return
	}
	goods, err := goodsSrv.GetGoodsDetail(goodsId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Msg: err,
		})
		return
	}
	c.JSON(http.StatusOK, Resp{
		Data: goods,
	})
}
