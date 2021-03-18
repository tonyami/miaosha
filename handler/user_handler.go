package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/service"
	"net/http"
)

type UserHandler struct {
	userService service.IUserService
}

func InitUserHandler() *UserHandler {
	return &UserHandler{userService: service.GetUserService()}
}

func (h *UserHandler) GetSmsCode(c *gin.Context) {
	mobile := c.Query("mobile")
	if len(mobile) != 11 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	code, err := h.userService.GetSmsCode(mobile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	}
}

func (h *UserHandler) UserLogin(c *gin.Context) {
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
	token, err := h.userService.Login(r.Mobile, r.SmsCode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	token, ok := c.Get("token")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userInfo, err := h.userService.GetUserInfo(token.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, userInfo)
	}
}
