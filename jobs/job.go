package jobs

func Init() {
	go GetOrderTimeoutJob().Start()
}
