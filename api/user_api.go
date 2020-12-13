package api

import (
	"github.com/gin-gonic/gin"
	"miaosha/lib/code"
	"miaosha/lib/resp"
	"strconv"
)

func getUserInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		resp.NewError(c, code.ConvertErr)
		return
	}
	user, err := userService.GetInfo(id)
	if err != nil {
		resp.NewError(c, err)
		return
	}
	resp.NewSuccessWith(c, user)
}
