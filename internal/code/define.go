package code

var (
	// public: [0, 1000]
	Success      = New(0, "成功")
	Unauthorized = New(401, "未授权")
	SystemErr    = New(555, "系统异常")

	// user: [1100, 1199]
	MobileErr = New(1100, "手机号码错误")
	CodeErr   = New(1101, "验证码错误")

	// goods: [1200, 1199]
	GoodsNotFound = New(1200, "商品不存在")

	// core: [1300, 1399]
	MiaoshaNotStarted = New(1300, "秒杀未开始")
	MiaoshaFinished   = New(1311, "秒杀已结束")
	MiaoshaSoldOut    = New(1322, "秒杀失败，已售罄")
	MiaoshaRepeated   = New(1333, "秒杀失败，重复秒杀")
	MiaoshaFailed     = New(1344, "秒杀失败")
)