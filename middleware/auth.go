package middleware

import (
	"github.com/gin-gonic/gin"
	"miaosha/service"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userService := service.GetUserService()
		user, err := userService.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if err = userService.RenewUserToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.Set("uid", user.Id)
		c.Set("mobile", user.Mobile)
		c.Set("token", token)
		c.Next()
	}
}
