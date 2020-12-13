package api

import (
	"github.com/gin-gonic/gin"
	"miaosha/lib/resp"
	"miaosha/model"
)

func userRegister(c *gin.Context) {
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	user := &model.User{
		Mobile:   mobile,
		Password: password,
	}
	if err := user.Check(); err != nil {
		resp.NewError(c, err)
		return
	}
	if err := userService.Register(user); err != nil {
		resp.NewError(c, err)
		return
	}
	resp.NewSuccess(c)
}
