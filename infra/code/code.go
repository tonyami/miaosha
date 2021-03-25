package code

var (
	// public [5000, 5100)
	DBErr           = code(5000, "系统错误:00")
	RedisErr        = code(5020, "系统错误:01")
	TooManyRequests = code(5040, "请求频繁")
	SerializeErr    = code(5050, "序列化错误")
	UnknownErr      = code(5099, "未知错误")

	// user [5100, 5200)
	GetSmsCodeErr = code(5100, "获取验证码失败")
	SmsCodeErr    = code(5101, "验证码错误")

	// goods [5200, 5300)
	GoodsSaleOut    = code(5200, "商品已售罄")
	MiaoshaNotStart = code(5211, "秒杀还未开始")
	MiaoshaEnded    = code(5222, "秒杀已结束")

	// order [5300, 5400)
	OrderNotFound  = code(5300, "订单不存在")
	OrderStatusErr = code(5310, "订单状态错误")
	OrderCloseErr  = code(5320, "订单取消失败")
	RepeateMiaosha = code(5330, "请勿重复秒杀")
	MiaoshaFailed  = code(5340, "秒杀失败")
)
