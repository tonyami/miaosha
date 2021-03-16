package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"miaosha/conf"
	"miaosha/infra/ip"
	"miaosha/infra/rdb"
	"net/http"
	"time"
)

const rateLimitKey = "rate_limit:%s"

var ctx = context.Background()

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := ip.Get(c.Request)
		count, err := rdb.Conn().Incr(ctx, fmt.Sprintf(rateLimitKey, clientIP)).Result()
		if err != nil {
			log.Printf("rdb.Incr() failed, err: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "系统错误",
			})
			return
		}
		if err = rdb.Conn().Expire(ctx, fmt.Sprintf(rateLimitKey, clientIP), time.Duration(conf.Conf.RateLimit.Time)*time.Second).Err(); err != nil {
			log.Printf("rdb.Expire() failed, err: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "系统错误",
			})
			return
		}
		if count > conf.Conf.RateLimit.Count {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"msg": "请求频繁",
			})
			return
		}
		c.Next()
	}
}
