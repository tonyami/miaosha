package conf

// redis
const (
	SmsCodeKey = "code:sms:%s" // 短信验证码
	SmsCodeIn  = 30            // 短信验证码有效期
	TokenKey   = "token:%s"    // 用户token
	TokenIn    = 1200          // 用户token有效期
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
	OrderUnPaid OrderStatus = 0 // 订单状态：未支付
	OrderPaid   OrderStatus = 1 // 订单状态：已支付
)
