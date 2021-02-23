package service

type OrderStatus int8 // 订单状态

const (
	Closed OrderStatus = -1 // 已关闭
	Unpaid OrderStatus = 0  // 待支付
	Paying OrderStatus = 1  // 支付中
	Paid   OrderStatus = 2  // 已支付
)

const (
	OrderExpire = 1800 // 订单有效期
	PageSize    = 10
)
