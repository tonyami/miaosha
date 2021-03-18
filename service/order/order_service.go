package order

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/infra/cache"
	"miaosha/infra/code"
	"miaosha/model"
	"miaosha/mq"
	"miaosha/repository"
	"miaosha/service"
	"sync"
)

var once sync.Once

func InitService() {
	once.Do(func() {
		service.OrderService = &orderService{
			orderRepository: repository.NewOrderRepository(),
			goodsService:    service.GetGoodsService(),
			redis:           cache.Client,
		}
	})
}

type orderService struct {
	orderRepository repository.OrderRepository
	goodsService    service.IGoodsService
	redis           *redis.Client
}

func (s *orderService) CountOrder(userId int64) (count model.OrderCount, err error) {
	if count, err = s.orderRepository.CountOrder(userId); err != nil {
		err = code.DBErr
	}
	return
}

func (s *orderService) GetOrderList(userId int64, status string, page int) (orderList []model.OrderInfoVO, err error) {
	var orderInfoList []model.OrderInfo
	if orderInfoList, err = s.orderRepository.GetOrderList(userId, status, page); err != nil {
		err = code.DBErr
		return
	}
	orderList = make([]model.OrderInfoVO, 0)
	for _, order := range orderInfoList {
		orderList = append(orderList, order.ToVO())
	}
	return
}

func (s *orderService) GetOrderInfoVO(orderId string, userId int64) (orderInfoVO model.OrderInfoVO, err error) {
	var orderInfo model.OrderInfo
	if orderInfo, err = s.GetOrderInfo(orderId); err != nil {
		err = code.DBErr
		return
	}
	if orderInfo.UserId != userId {
		err = code.OrderNotFound
		return
	}
	orderInfoVO = orderInfo.ToVO()
	return
}

func (s *orderService) Miaosha(userId, goodsId int64) (err error) {
	var (
		orderId string
		stock   int64
	)
	// 校验重复秒杀
	if orderId, err = s.GetOrderId(userId, goodsId); err != nil {
		err = code.RedisErr
		return
	}
	if len(orderId) > 0 {
		err = code.RepeateMiaosha
		return
	}
	// 预减库存
	if stock, err = s.goodsService.DecrStock(goodsId); err != nil {
		err = code.RedisErr
		return
	}
	if stock < 0 {
		err = code.GoodsSaleOut
		return
	}
	// 异步下单
	msg := mq.PrecreateOrderMsg{
		UserId:  userId,
		GoodsId: goodsId,
	}
	if err = mq.PrecreateOrder.Send(msg); err != nil {
		err = code.RedisErr
	}
	return
}

func (s *orderService) GetMiaoshaReuslt(userId, goodsId int64) (result model.MiaoshaResult, err error) {
	var orderId string
	if orderId, err = s.GetOrderId(userId, goodsId); err != nil {
		err = code.RedisErr
		return
	}
	if len(orderId) == 0 {
		result = model.MiaoshaResult{Status: 0} // 排队中
	} else {
		result = model.MiaoshaResult{Status: 1, OrderId: orderId} // 秒杀成功
	}
	return
}

func (s *orderService) CloseOrder(userId int64, orderId string) (err error) {
	var orderInfo model.OrderInfo
	if orderInfo, err = s.orderRepository.GetOrderInfo(orderId); err != nil {
		err = code.DBErr
		return
	}
	if orderInfo.Id == 0 || orderInfo.UserId != userId {
		err = code.OrderNotFound
		return
	}
	if orderInfo.Status != model.Unpaid {
		err = code.OrderStatusErr
		return
	}
	if err = s.orderRepository.CloseOrder(orderInfo); err != nil {
		err = code.OrderCloseErr
		return
	}
	// 加缓存中库存
	if err = s.goodsService.IncrStock(orderInfo.GoodsId); err != nil {
		err = code.RedisErr
		return
	}
	// 删除订单缓存
	if err = s.DeleteOrderCache(orderInfo); err != nil {
		err = code.RedisErr
		return
	}
	// 移除延迟队列中订单号
	mq.OrderTimeout.Remove(orderInfo.OrderId)
	return
}

func (s *orderService) GetOrderInfo(orderId string) (orderInfo model.OrderInfo, err error) {
	if orderInfo, err = s.orderRepository.GetOrderInfo(orderId); err != nil {
		err = code.DBErr
	}
	return
}

func (s *orderService) CreateOrder(userId, goodsId int64) (err error) {
	var (
		orderId string
		goods   model.Goods
	)
	// 是否重复秒杀
	if orderId, err = s.GetOrderId(userId, goodsId); err != nil {
		log.Printf("s.GetOrderId() failed, err: %v", err)
		return
	}
	if len(orderId) > 0 {
		return
	}
	// 查询商品信息
	if goods, err = s.goodsService.GetGoods(goodsId); err != nil {
		log.Printf("goodsService.GetGoods() failed, err: %v", err)
		return
	}
	// 创建订单
	orderInfo := model.NewOrderInfo(userId, goods)
	if err = s.orderRepository.CreateOrder(orderInfo); err != nil {
		log.Printf("orderRepository.CreateOrder() failed, err: %v", err)
		return
	}
	// 创建订单信息缓存
	if err = s.CreateOrderCache(orderInfo); err != nil {
		return
	}
	// 加入订单超时延迟队列
	mq.OrderTimeout.Send(orderInfo.OrderId)
	return
}

func (s *orderService) CreateOrderCache(order model.OrderInfo) (err error) {
	if err = s.redis.Set(service.Ctx, fmt.Sprintf(service.OrderUidGidKey, order.UserId, order.GoodsId), order.OrderId, -1).Err(); err != nil {
		log.Printf("redis.Set() failed, err: %v, order: %v", err, order)
		err = code.RedisErr
	}
	return
}

func (s *orderService) DeleteOrderCache(order model.OrderInfo) (err error) {
	if err = s.redis.Del(service.Ctx, fmt.Sprintf(service.OrderUidGidKey, order.UserId, order.GoodsId)).Err(); err != nil {
		log.Printf("redis.Del() falied, err: %v, order: %v", err, order)
		err = code.RedisErr
	}
	return
}

func (s *orderService) GetOrderId(userId, goodsId int64) (orderId string, err error) {
	if orderId, err = s.redis.Get(service.Ctx, fmt.Sprintf(service.OrderUidGidKey, userId, goodsId)).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("redis.Get() failed, err: %v", err)
			err = code.RedisErr
		}
	}
	return
}
