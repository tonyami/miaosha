package api

func initRouter() {
	e.POST("/user/register", userRegister)
}
