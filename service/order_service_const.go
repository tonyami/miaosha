package service

type OrderStatus int8 // 订单状态

const (
	Closed OrderStatus = -1 // 已关闭
	Unpaid OrderStatus = 0  // 待支付
	Paying OrderStatus = 1  // 支付中
	Paid   OrderStatus = 2  // 已支付
)

type OrderListStatus string

const (
	All        OrderListStatus = "all"
	Unfinished OrderListStatus = "unfinished"
	Finished   OrderListStatus = "finished"
	Closed2    OrderListStatus = "closed"
)

func OrderListStatusContains(status string) (OrderListStatus, bool) {
	list := []OrderListStatus{All, Unfinished, Finished, Closed2}
	for _, v := range list {
		if status == string(v) {
			return v, true
		}
	}
	return "", false
}

const (
	OrderTimeout = 60 // 订单有效期
	PageSize     = 10
)
