package handler

import (
	"github.com/gin-gonic/gin"
	"miaosha/infra/key"
	"miaosha/model"
	"miaosha/repository"
	"net/http"
)

const (
	SystemErr     = "系统错误"
	SmsCodeSize   = 6
	UserTokenSize = 64
)

func GetSmsCode(c *gin.Context) {
	mobile := c.Query("mobile")
	if len(mobile) != 11 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// 生成短信验证码并保存
	code := key.Create(key.Number, SmsCodeSize)
	if err := repository.SaveLoginSmsCode(mobile, code); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "获取验证码失败",
		})
		return
	}
	// 模拟短信发送成功
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
	return
}

func UserLogin(c *gin.Context) {
	r := new(struct {
		Mobile  string `form:"mobile" binding:"required"`
		SmsCode string `form:"smsCode" binding:"required"`
	})
	// 参数绑定及简单校验
	if err := c.Bind(r); err != nil {
		return
	}
	if len(r.Mobile) != 11 || len(r.SmsCode) != 6 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// 获取验证码并比对
	code, err := repository.GetLoginSmsCode(r.Mobile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	if code != r.SmsCode {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}
	// 根据手机号码查询用户
	var user model.User
	if user, err = repository.GetUserByMobile(r.Mobile); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	// 用户不存在则自动注册
	if user.Id == 0 {
		user = model.User{
			Mobile: r.Mobile,
		}
		if user.Id, err = repository.SaveUser(user); err != nil || user.Id == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": SystemErr,
			})
			return
		}
	}
	// 生成token并保存
	token := key.Create(key.AlphabetAndNumber, UserTokenSize)
	if err = repository.SaveUserToken(token, user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	// 删除验证码
	if err = repository.DeleteLoginSmsCode(r.Mobile); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
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
	user, err := repository.GetUserByToken(token.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	count, err := repository.CountOrder(user.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": SystemErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"count": count,
	})
}
