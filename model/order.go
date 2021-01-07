package model

import (
	"miaosha/conf"
	"time"
)

type Order struct {
	Id         string           `json:"id"`
	UserId     int64            `json:"userId"`
	GoodsId    int64            `json:"goodsId"`
	GoodsName  string           `json:"goodsName"`
	GoodsImg   string           `json:"goodsImg"`
	Price      int64            `json:"price"`
	CreateTime time.Time        `json:"createTime"`
	Status     conf.OrderStatus `json:"status"`
	PayTime    time.Time        `json:"payTime,omitempty"`
}
