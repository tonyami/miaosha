package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := userService.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}
		if user == nil || user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.Abort()
			return
		}
		if err = userService.RenewToken(token); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("token", token)
		c.Set("uid", user.Id)
		c.Set("mobile", user.Mobile)
		c.Next()
	}
}
