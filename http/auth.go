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
			json2(c, nil, code.Unauthorized)
			c.Abort()
			return
		}
		user, err := userService.Auth(token)
		if err != nil {
			json2(c, nil, err)
			c.Abort()
			return
		}
		if user == nil {
			json2(c, nil, code.Unauthorized)
			c.Abort()
			return
		}
		c.Set(conf.User, user)
		c.Next()
	}
}
