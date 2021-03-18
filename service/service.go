package service

import (
	"miaosha/model"
)

var (
	UserService  IUserService
	GoodsService IGoodsService
	OrderService IOrderService
)

func GetUserService() IUserService {
	return UserService
}

func GetGoodsService() IGoodsService {
	return GoodsService
}

func GetOrderService() IOrderService {
	return OrderService
}

type IUserService interface {
	SaveLoginSmsCode(mobile, code string) (err error)
	GetLoginSmsCode(mobile string) (code string, err error)
	DeleteLoginSmsCode(mobile string) (err error)
	SaveUserToken(token string, user model.User) (err error)
	GetUserByToken(token string) (user model.User, err error)
	RenewUserToken(token string) (err error)
	GetSmsCode(string) (string, error)
	Login(string, string) (string, error)
	GetUserInfo(string) (model.UserInfoVO, error)
}

type IGoodsService interface {
	InitGoodsStock() error
	GetGoodsList(int) ([]model.GoodsVO, error)
	GetGoodsVO(int64) (model.GoodsVO, error)
	GetGoods(int64) (model.Goods, error)
	SetGoodsStock(goodsId int64, stock int) (err error)
	DecrStock(goodsId int64) (stock int64, err error)
	IncrStock(goodsId int64) (err error)
}

type IOrderService interface {
	GetOrderList(int64, string, int) ([]model.OrderInfoVO, error)
	GetOrderInfoVO(string, int64) (model.OrderInfoVO, error)
	Miaosha(int64, int64) (err error)
	GetMiaoshaReuslt(int64, int64) (model.MiaoshaResult, error)
	CloseOrder(int64, string) error
	CountOrder(int64) (model.OrderCount, error)
	GetOrderInfo(string) (model.OrderInfo, error)
	CreateOrder(int64, int64) error
	CreateOrderCache(model.OrderInfo) (err error)
	DeleteOrderCache(model.OrderInfo) (err error)
	GetOrderId(int64, int64) (string, error)
}
