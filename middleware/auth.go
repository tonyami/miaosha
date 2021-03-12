package middleware

import (
	"github.com/gin-gonic/gin"
	"miaosha/repository"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := repository.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if err = repository.RenewUserToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "系统错误",
			})
			return
		}
		c.Set("uid", user.Id)
		c.Set("mobile", user.Mobile)
		c.Set("token", token)
		c.Next()
	}
}
