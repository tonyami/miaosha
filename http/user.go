package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/internal/code"
)

func SendSmsCode(c *gin.Context) {
	r := new(struct {
		Mobile string `form:"mobile" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	if len(r.Mobile) != 11 {
		JSON2(c, nil, code.MobileErr)
		return
	}
	smsCode, err := userService.SendSmsCode(r.Mobile)
	JSON2(c, smsCode, err)
}

func Login(c *gin.Context) {
	r := new(struct {
		Mobile  string `form:"mobile" binding:"required"`
		SmsCode string `form:"smsCode" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	if len(r.Mobile) != 11 {
		JSON2(c, nil, code.MobileErr)
		return
	}
	if len(r.SmsCode) != 6 {
		JSON2(c, nil, code.CodeErr)
		return
	}
	token, err := userService.Login(r.Mobile, r.SmsCode)
	JSON2(c, token, err)
}

func GetUser(c *gin.Context) {
	user, ok := c.Get(conf.User)
	if !ok {
		JSON2(c, nil, code.Unauthorized)
		return
	}
	JSON2(c, user, nil)
}
