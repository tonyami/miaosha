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
		json2(c, nil, code.MobileErr)
		return
	}
	smsCode, err := userService.SendSmsCode(r.Mobile)
	json2(c, smsCode, err)
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
		json2(c, nil, code.MobileErr)
		return
	}
	if len(r.SmsCode) != 6 {
		json2(c, nil, code.CodeErr)
		return
	}
	token, err := userService.Login(r.Mobile, r.SmsCode)
	json2(c, token, err)
}

func GetUser(c *gin.Context) {
	user, ok := c.Get(conf.User)
	if !ok {
		json2(c, nil, code.Unauthorized)
		return
	}
	json2(c, user, nil)
}
