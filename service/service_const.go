package service

import "context"

const (
	SmsCodeSize        = 6
	UserTokenSize      = 64
	LoginSmsCodeKey    = "login:sms_code:%s"
	LoginSmsCodeExpire = 30
	UserTokenKey       = "user:token:%s"
	UserTokenExpire    = 3600
	GoodsStockKey      = "goods_stock:%d"
	OrderUidGidKey     = "order:%d:%d"
)

var Ctx = context.Background()
