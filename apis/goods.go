package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetGoodsList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	list, err := goodsService.GetList(page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetGoods(c *gin.Context) {
	goodsId, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	goods, err := goodsService.Get(goodsId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, goods)
}
