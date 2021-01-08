package model

import (
	"miaosha/conf"
	"time"
)

type Order struct {
	Id         string
	UserId     int64
	GoodsId    int64
	Status     conf.OrderStatus
	CreateTime time.Time
}
