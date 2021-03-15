package mq

func Init() {
	go OrderTimeout.Receive()
	go OrderPrecreate.Receive()
}
