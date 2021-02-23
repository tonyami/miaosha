package service

import (
	"time"
)

var IOrderService OrderService

func GetOrderService() OrderService {
	return IOrderService
}

type OrderService interface {
	GetList(int64, int, string) ([]*OrderDTO, error)      // 获取订单列表
	GetById(int64) (*OrderDTO, error)                     // 获取订单详情
	GetByIdAndUserId(id, userId int64) (*OrderDTO, error) // 获取订单详情
	Create(int64, int64) (int64, error)                   // 创建订单
	ManualCancel(int64, int64) error                      // 手动取消订单
	AutoCancel(int64) error                               // 自动取消订单
}

type OrderDTO struct {
	Id         int64       `json:"id"`
	GoodsId    int64       `json:"goodsId"`
	GoodsName  string      `json:"goodsName"`
	GoodsImg   string      `json:"goodsImg"`
	GoodsPrice int64       `json:"goodsPrice"`
	Status     OrderStatus `json:"status"`
	Duration   int64       `json:"duration,omitempty"`
	CreateTime time.Time   `json:"createTime"`
	UpdateTime time.Time   `json:"updateTime"`
}
