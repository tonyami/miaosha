package api

func initRouter() {
	e.GET("/user/info", getUserInfo)
}
