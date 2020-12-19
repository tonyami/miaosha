package middleware

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/service/user"
	"net/http"
)

func Auth(userSrv *user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		dto, err := userSrv.Info(token)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if dto.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(conf.UserSession, dto)
		c.Next()
	}
}
