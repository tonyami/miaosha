package code

var (
	// public: [0, 1000]
	Success      = New(0, "成功")
	Unauthorized = New(111, "未授权")
	Denied       = New(222, "无权限")
	ParseErr     = New(333, "解析失败")
	SystemErr    = New(999, "系统异常")

	// user: [1100, 1199]
	MobileErr = New(1100, "手机号码错误")
	CodeErr   = New(1111, "验证码错误")

	// goods: [1200, 1199]
	GoodsNotFound = New(1200, "商品不存在")

	// order: [1300, 1399]
	MiaoshaNotStarted = New(1300, "秒杀未开始")
	MiaoshaFinished   = New(1311, "秒杀已结束")
	MiaoshaSoldOut    = New(1322, "已售罄")
	MiaoshaRepeated   = New(1333, "请勿重复秒杀")
	MiaoshaFailed     = New(1344, "秒杀失败")
	OrderNotFound     = New(1355, "订单不存在")
	OrderCannotClose  = New(1366, "订单取消失败：状态错误")
	OrderCloseFailed  = New(1377, "订单取消失败")

	// pay: [1400, 1499]
	PayRequestErr     = New(1400, "支付请求失败")
	PayOrderStatusErr = New(1401, "支付失败，订单状态错误")
)
