package order

import (
	"miaosha/conf"
	"miaosha/service"
	"time"
)

type Order struct {
	Id         int64
	UserId     int64
	GoodsId    int64
	GoodsName  string
	GoodsImg   string
	GoodsPrice int64
	PaymentId  int64
	Status     int8
	CreateTime time.Time
	UpdateTime time.Time
}

func (o *Order) toDTO() *service.OrderDTO {
	c := conf.Conf.Order
	dto := &service.OrderDTO{
		Id:         o.Id,
		GoodsId:    o.GoodsId,
		GoodsName:  o.GoodsName,
		GoodsImg:   o.GoodsImg,
		GoodsPrice: o.GoodsPrice,
		Status:     service.OrderStatus(o.Status),
		Duration:   0,
		Timeout:    c.Expire,
		CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: o.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if dto.Status == service.Unpaid {
		dto.Duration = o.CreateTime.Unix() + c.Expire - time.Now().Unix()
	}
	return dto
}

type OrderCount struct {
	Unfinished int64
	Finished   int64
	Closed     int64
}

func (o *OrderCount) toDTO() *service.OrderCountDTO {
	return &service.OrderCountDTO{
		Unfinished: o.Unfinished,
		Finished:   o.Finished,
		Closed:     o.Closed,
	}
}
