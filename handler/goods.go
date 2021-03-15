package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/model"
	"miaosha/repository"
	"net/http"
	"strconv"
)

func ReloadGoodsStock(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	list, err := repository.GetGoodsList(page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	for _, v := range list {
		if err = repository.SetGoodsStock(v.Id, v.Stock); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": SystemErr,
			})
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}

func GetGoodsList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	list, err := repository.GetGoodsList(page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	goodsList := make([]model.GoodsVO, 0)
	for _, v := range list {
		goodsList = append(goodsList, v.ToVO())
	}
	c.JSON(http.StatusOK, goodsList)
}

func GetGoods(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	goods, err := repository.GetGoods(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	c.JSON(http.StatusOK, goods.ToVO())
}
