package order

import (
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"miaosha/dao/order"
	"miaosha/internal/cache"
	"miaosha/internal/code"
	"miaosha/model"
	"miaosha/service/goods"
	"miaosha/service/user"
	"miaosha/util/key"
	"time"
)

type Service struct {
	userService  *user.Service
	goodsService *goods.Service
	redisCli     *redis.Client
	dao          *order.Dao
}

func New(goodsService *goods.Service) *Service {
	return &Service{
		dao:          order.New(),
		goodsService: goodsService,
		redisCli:     cache.New(conf.Conf.Redis),
	}
}

// 核心功能：秒杀
func (s *Service) Miaosha(userId int64, goodsId int64) (order *model.Order, err error) {
	var g *model.Goods
	// 1、查询秒杀商品
	if g, err = s.goodsService.GetGoods(goodsId); err != nil {
		return
	}
	// 2、校验秒杀开始时间、结束时间、库存
	if err = check(g); err != nil {
		return
	}
	// 3、校验是否已经秒杀过
	var count int64
	if count, err = s.dao.Count(userId, g.Id); err != nil {
		log.Printf("Miaosha Failed: %s", err)
		err = code.SystemErr
		return
	}
	if count > 0 {
		err = code.MiaoshaRepeated
		return
	}
	// 4、减库存、生成订单
	order = &model.Order{
		Id:         key.CreateOrderId(),
		UserId:     userId,
		GoodsId:    g.Id,
		GoodsName:  g.Name,
		GoodsImg:   g.Img,
		Price:      g.Price,
		CreateTime: time.Now(),
		Status:     conf.OrderPayWaiting,
	}
	if err = s.dao.Miaosha(order); err != nil {
		err = code.MiaoshaFailed
		order = nil
	}
	return
}

// 校验秒杀开始时间、结束时间、库存
func check(goods *model.Goods) (err error) {
	now := time.Now().Unix()
	startTime := goods.StartTime.Unix()
	endTime := goods.EndTime.Unix()
	if now < startTime {
		err = code.MiaoshaNotStarted
	} else if now > endTime {
		err = code.MiaoshaFinished
	} else if goods.Stock <= 0 {
		err = code.MiaoshaSoldOut
	}
	return
}
