package order

import (
	"log"
	"miaosha/conf"
	"miaosha/dao/order"
	"miaosha/internal/code"
	"miaosha/internal/key"
	"miaosha/model"
	"miaosha/service/goods"
	"miaosha/service/user"
	"time"
)

type Service struct {
	userService  *user.Service
	goodsService *goods.Service
	dao          *order.Dao
}

func New(goodsService *goods.Service) *Service {
	return &Service{
		dao:          order.New(),
		goodsService: goodsService,
	}
}

func (s *Service) Get(orderId string, userId int64) (order *model.OrderDTO, err error) {
	if order, err = s.dao.Get(orderId); err != nil {
		log.Printf("【Order】Get Failed: %s", err)
		err = code.SystemErr
	}
	if order == nil {
		err = code.OrderNotFound
	}
	// 禁止访问非本人订单
	if order.UserId != userId {
		err = code.Denied
	}
	return
}

func (s *Service) GetList(userId int64, page int, status string) (orders []*model.OrderDTO, err error) {
	if orders, err = s.dao.GetList(userId, page, conf.PageSize, status); err != nil {
		log.Printf("【Order】GetList Failed: %s", err)
		err = code.SystemErr
	}
	return
}

// 核心功能：秒杀
func (s *Service) Miaosha(userId int64, goodsId int64) (orderId string, err error) {
	var g *model.Goods
	// 1、查询秒杀商品
	if g, err = s.goodsService.Get(goodsId); err != nil {
		return
	}
	// 2、校验秒杀开始时间、结束时间、库存
	if err = g.Check(); err != nil {
		return
	}
	// 3、校验是否已经秒杀过
	var count int64
	if count, err = s.dao.Count(userId, g.Id); err != nil {
		log.Printf("【Order】Miaosha Failed: %s", err)
		err = code.SystemErr
		return
	}
	if count > 0 {
		err = code.MiaoshaRepeated
		return
	}
	// 4、减库存、生成订单
	orderId = key.OrderId()
	o := &model.Order{
		Id:         orderId,
		UserId:     userId,
		GoodsId:    g.Id,
		CreateTime: time.Now(),
		Status:     conf.OrderUnPaid,
	}
	if err = s.dao.Miaosha(o); err != nil {
		err = code.MiaoshaFailed
		orderId = ""
	}
	return
}
