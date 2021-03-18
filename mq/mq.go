package mq

import (
	"context"
	"miaosha/infra/cache"
	"miaosha/service"
)

var ctx = context.Background()

func Init() {
	OrderTimeout = &orderTimeout{
		orderService: service.GetOrderService(),
		redis:        cache.Client,
	}
	PrecreateOrder = &precreateOrder{
		orderService: service.GetOrderService(),
		redis:        cache.Client,
	}
}

func Run() {
	go OrderTimeout.Receive()
	go PrecreateOrder.Receive()
}
