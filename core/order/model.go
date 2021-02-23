package order

import (
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
	dto := &service.OrderDTO{
		Id:         o.Id,
		GoodsId:    o.GoodsId,
		GoodsName:  o.GoodsName,
		GoodsImg:   o.GoodsImg,
		GoodsPrice: o.GoodsPrice,
		Status:     service.OrderStatus(o.Status),
		Duration:   0,
		CreateTime: o.CreateTime,
		UpdateTime: o.UpdateTime,
	}
	if dto.Status == service.Unpaid {
		dto.Duration = dto.CreateTime.Unix() + service.OrderExpire - time.Now().Unix()
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
