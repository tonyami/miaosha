package errmsg

var (
	InvalidParameter Errmsg = "请求参数错误"
	SystemErr        Errmsg = "系统错误"

	MobileRegistered    Errmsg = "注册失败，该手机号码已注册"
	RegisterErr         Errmsg = "注册失败"
	MobileNotRegistered Errmsg = "登录失败，该手机号码未注册"
	PasswordErr         Errmsg = "登录失败，密码错误"
)

type Errmsg string

func (err Errmsg) Error() string {
	return string(err)
}
