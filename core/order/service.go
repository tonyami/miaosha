package order

import (
	"database/sql"
	"errors"
	"log"
	"miaosha/core/goods"
	"miaosha/internal/db"
	"miaosha/service"
	"sync"
)

var once sync.Once

func init() {
	log.Printf("init order service...")
	once.Do(func() {
		service.IOrderService = &orderService{
			userService:  service.GetUserService(),
			goodsService: service.GetGoodsService(),
		}
	})
}

type orderService struct {
	userService  service.UserService
	goodsService service.GoodsService
}

func (s *orderService) GetList(userId int64, page int, status string) (orders []*service.OrderDTO, err error) {
	dao := NewDao(db.Get())
	var list []*Order
	if list, err = dao.GetList(userId, page, service.PageSize, status); err != nil {
		err = errors.New("db error")
		return
	}
	orders = []*service.OrderDTO{}
	for _, item := range list {
		orders = append(orders, item.toDTO())
	}
	return
}

func (s *orderService) Get(id, userId int64) (order *service.OrderDTO, err error) {
	dao := NewDao(db.Get())
	var o *Order
	if o, err = dao.Get(id, userId); err != nil {
		err = errors.New("db error")
		return
	}
	if o.Id == 0 {
		err = errors.New("goods not found")
		return
	}
	order = o.toDTO()
	return
}

func countPurchased(userId, goodsId int64) (count int64, err error) {
	conn := db.Get()
	dao := NewDao(conn)
	if count, err = dao.CountPurchased(userId, goodsId); err != nil {
		err = errors.New("db error")
		return
	}
	return
}

func (s *orderService) Create(userId, goodsId int64) (orderId int64, err error) {
	var dto *service.GoodsDTO
	var count int64
	// 1、查询秒杀商品
	if dto, err = s.goodsService.Get(goodsId); err != nil {
		return
	}
	// 2、校验秒杀开始时间、结束时间、库存
	if err = dto.Check(); err != nil {
		return
	}
	// 3、校验是否重复秒杀
	if count, err = countPurchased(userId, goodsId); err != nil {
		return
	}
	if count > 0 {
		err = errors.New("请勿重复秒杀")
		return
	}
	// 4、生成订单、减库存
	order := &Order{
		UserId:     userId,
		GoodsId:    dto.Id,
		GoodsName:  dto.Name,
		GoodsImg:   dto.Img,
		GoodsPrice: dto.Price,
	}
	conn := db.Get()
	orderDao := NewDao(conn)
	goodsDao := goods.NewDao(conn)
	if err = db.Tx(conn, func(conn *sql.DB) (err error) {
		// 生成订单
		if orderId, err = orderDao.Insert(order); err != nil {
			return
		}
		if orderId == 0 {
			err = errors.New("create order failed")
			return
		}
		// 减库存
		if err = goodsDao.DecrStock(dto.Id); err != nil {
			return
		}
		return
	}); err != nil {
		err = errors.New("秒杀失败")
		return
	}
	return
}

func (s *orderService) Cancel(id, userId int64) (err error) {
	var dto *service.OrderDTO
	if dto, err = s.Get(id, userId); err != nil {
		return
	}
	if dto.Status != service.Unpaid {
		err = errors.New("订单无法取消")
		return
	}
	conn := db.Get()
	orderDao := NewDao(conn)
	goodsDao := goods.NewDao(conn)
	if err = db.Tx(conn, func(conn *sql.DB) (err error) {
		// 关闭订单
		if err = orderDao.Close(dto.Id); err != nil {
			return
		}
		// 加库存
		if err = goodsDao.IncrStock(dto.GoodsId); err != nil {
			return
		}
		return
	}); err != nil {
		err = errors.New("订单取消失败")
		return
	}
	return
}
