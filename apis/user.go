package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSmsCode(c *gin.Context) {
	r := new(struct {
		Mobile string `form:"mobile" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if len(r.Mobile) != 11 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	code, err := userService.SendSmsCode(r.Mobile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

func Login(c *gin.Context) {
	r := new(struct {
		Mobile  string `form:"mobile" binding:"required"`
		SmsCode string `form:"smsCode" binding:"required"`
	})
	if err := c.Bind(r); err != nil {
		return
	}
	if len(r.Mobile) != 11 || len(r.SmsCode) != 6 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	token, err := userService.Login(r.Mobile, r.SmsCode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetUserInfo(c *gin.Context) {
	token, ok := c.Get("token")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, err := userService.GetUserByToken(token.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
