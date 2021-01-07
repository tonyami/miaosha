package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/internal/code"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			JSON2(c, nil, code.Unauthorized)
			c.Abort()
			return
		}
		user, err := userService.GetUser(token)
		if err != nil {
			JSON2(c, nil, err)
			c.Abort()
			return
		}
		if user == nil {
			JSON2(c, nil, code.Unauthorized)
			c.Abort()
			return
		}
		c.Set(conf.UserSession, user)
		c.Next()
	}
}
