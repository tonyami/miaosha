package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"miaosha/conf"
	"miaosha/infra/cache"
	"miaosha/infra/code"
	"miaosha/infra/util"
	"net/http"
	"time"
)

const rateLimitKey = "rate_limit:%s"

var ctx = context.Background()

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := util.GetIP(c.Request)
		count, err := cache.Client.Incr(ctx, fmt.Sprintf(rateLimitKey, clientIP)).Result()
		if err != nil {
			log.Printf("rdb.Incr() failed, err: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": code.RedisErr.Error(),
			})
			return
		}
		if err = cache.Client.Expire(ctx, fmt.Sprintf(rateLimitKey, clientIP), time.Duration(conf.Conf.RateLimit.Time)*time.Second).Err(); err != nil {
			log.Printf("rdb.Expire() failed, err: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": code.RedisErr.Error(),
			})
			return
		}
		if count > conf.Conf.RateLimit.Count {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"msg": code.TooManyRequests.Error(),
			})
			return
		}
		c.Next()
	}
}
