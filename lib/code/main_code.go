package code

var (
	// system
	Success           = code(0, "成功")
	InvalidParameters = code(50000, "请求参数错误")
	NotAuthorization  = code(50011, "无权限")
	ConvertErr        = code(50022, "转换失败")
	SerializeErr      = code(50033, "序列划失败")
	NotAuthentication = code(50044, "未认证，请先登录")
	SystemErr         = code(50055, "系统错误")

	// user
	UserNotFound = code(40001, "未找到该用户")
)
