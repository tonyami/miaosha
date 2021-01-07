package conf

import (
	"time"
)

// redis
const (
	SmsCodeKey = "code:sms:%s"    // 短信验证码
	SmsCodeIn  = 60 * time.Second // 短信验证码有效期
	TokenKey   = "token:%s"       // 用户token
	TokenIn    = 20 * time.Minute // 用户token有效期
)

const (
	User     = "X-User"
	PageSize = 10 // 商品列表分页大小
)

type MiaoshaStatus int8

const (
	MiaoshaNotStarted MiaoshaStatus = 0 // 秒杀状态：未开始
	MiaoshaOnGoing    MiaoshaStatus = 1 // 秒杀状态：进行中
	MiaoshaSoldOut    MiaoshaStatus = 2 // 秒杀状态：售罄
	MiaoshaFinished   MiaoshaStatus = 3 // 秒杀状态：已结束
)

type OrderStatus int8

const (
	OrderPayWaiting OrderStatus = 1 // 订单状态：秒杀成功，待支付
	OrderPaySuccess OrderStatus = 2 // 订单状态：支付成功
)
