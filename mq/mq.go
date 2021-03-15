package mq

func Init() {
	go OrderTimeout.Receive()
}
