package model

import (
	"miaosha/conf"
	"miaosha/infra/key"
	"time"
)

const (
	Closed int8 = -1
	Unpaid int8 = 0
	Paying int8 = 1
	Paid   int8 = 2
)

type Order struct {
	Id         int64     `db:"id"`
	OrderId    string    `db:"order_id"`
	UserId     int64     `db:"user_id"`
	GoodsId    int64     `db:"goods_id"`
	GoodsName  string    `db:"goods_name"`
	GoodsImg   string    `db:"goods_img"`
	GoodsPrice int64     `db:"goods_price"`
	PaymentId  int64     `db:"payment_id"`
	Status     int8      `db:"status"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

type OrderVO struct {
	OrderId    string `json:"orderId"`
	GoodsId    int64  `json:"goodsId"`
	GoodsName  string `json:"goodsName"`
	GoodsImg   string `json:"goodsImg"`
	GoodsPrice int64  `json:"goodsPrice"`
	Status     int8   `json:"status"`
	Duration   int64  `json:"duration"`
	Timeout    int64  `json:"timeout"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func NewOrder(userId int64, goods Goods) Order {
	order := Order{
		OrderId:    createOrderId(),
		UserId:     userId,
		GoodsId:    goods.Id,
		GoodsName:  goods.Name,
		GoodsImg:   goods.Img,
		GoodsPrice: goods.Price,
		Status:     Unpaid,
	}
	return order
}

func createOrderId() string {
	return time.Now().Format("20060102150405") + key.Create(key.Number, 6)
}

func (order Order) ToVO() OrderVO {
	c := conf.Conf.Order
	orderVO := OrderVO{
		OrderId:    order.OrderId,
		GoodsId:    order.GoodsId,
		GoodsName:  order.GoodsName,
		GoodsImg:   order.GoodsImg,
		GoodsPrice: order.GoodsPrice,
		Status:     order.Status,
		Duration:   0,
		Timeout:    c.Expire,
		CreateTime: order.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: order.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if orderVO.Status == 0 {
		orderVO.Duration = order.CreateTime.Unix() + c.Expire - time.Now().Unix()
	}
	return orderVO
}

type OrderCount struct {
	Unfinished int64 `db:"unfinished" json:"unfinished"`
	Finished   int64 `db:"finished" json:"finished"`
	Closed     int64 `db:"closed" json:"closed"`
}
