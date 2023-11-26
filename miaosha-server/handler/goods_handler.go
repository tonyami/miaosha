package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/service"
	"net/http"
	"strconv"
)

type GoodsHandler struct {
	goodsService service.IGoodsService
}

func InitGoodsHandler() *GoodsHandler {
	return &GoodsHandler{goodsService: service.GetGoodsService()}
}

func (h *GoodsHandler) InitGoodsStock(c *gin.Context) {
	if err := h.goodsService.InitGoodsStock(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func (h *GoodsHandler) GetGoodsList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	list, err := h.goodsService.GetGoodsList(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func (h *GoodsHandler) GetGoods(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	goods, err := h.goodsService.GetGoodsVO(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, goods)
	}
}
