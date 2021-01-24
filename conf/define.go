package conf

// redis
const (
	SmsCodeKey = "code:sms:%s" // 短信验证码
	SmsCodeIn  = 30            // 短信验证码有效期
	TokenKey   = "token:%s"    // 用户token
	TokenIn    = 1200          // 用户token有效期
)

const (
	User                   = "X-User"
	PageSize               = 10  // 默认分页大小
	OrderExpire            = 600 // 订单有效期，过期未付自动关闭
	OrderSchedulerInterval = 30  // 超时订单扫描间隔
)

type MiaoshaStatus int8

const (
	MiaoshaFinished   MiaoshaStatus = -1 // 秒杀状态：已结束
	MiaoshaNotStarted MiaoshaStatus = 0  // 秒杀状态：未开始
	MiaoshaOnGoing    MiaoshaStatus = 1  // 秒杀状态：进行中
	MiaoshaSoldOut    MiaoshaStatus = 2  // 秒杀状态：售罄
)

type OrderStatus int8

const (
	OrderStatusClosed     OrderStatus = -1 // 订单状态：已关闭
	OrderStatusUnfinished OrderStatus = 0  // 订单状态：未完成
	OrderStatusFinished   OrderStatus = 1  // 订单状态：已完成
)
