package code

var (
	// system [50000, 50099]
	Success           = code(0, "成功")
	InvalidParams     = code(50000, "请求参数错误")
	NotAuthentication = code(50011, "未登录，请登录后再试")
	SystemErr         = code(50055, "系统错误")

	// user [50100, 50199]
	UserRegisterErr      = code(50100, "注册失败")
	UserMobileRegistered = code(50111, "注册失败，手机号码已注册")
)
