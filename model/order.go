package model

import (
	"miaosha/conf"
	"time"
)

type OrderDTO struct {
	Id         string           `json:"id"`
	UserId     int64            `json:"-"`
	GoodsId    int64            `json:"gid"`
	Name       string           `json:"name"`
	Img        string           `json:"img"`
	Price      int64            `json:"price"`
	Status     conf.OrderStatus `json:"status"`
	CreateTime time.Time        `json:"createTime"`
}

type Order struct {
	Id         string
	UserId     int64
	GoodsId    int64
	Status     conf.OrderStatus
	CreateTime time.Time
}
