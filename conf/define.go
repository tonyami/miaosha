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
	DefaultAvatar    = "default.jpg"
	UserSession      = "X-User-Session" // user session
	PageSize         = 10               // 商品列表分页大小
	StatusNotStarted = 0                // 秒杀状态：未开始
	StatusOnGoing    = 1                // 秒杀状态：进行中
	StatusSoldOut    = 2                // 秒杀状态：售罄
	StatusFinished   = 3                // 秒杀状态：已结束
)
