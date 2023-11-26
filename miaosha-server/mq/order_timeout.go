package mq

import (
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"miaosha/model"
	"miaosha/service"
	"strconv"
	"time"
)

const orderTimeoutDelayQueue = "order_timeout_delay_queue"

var OrderTimeout *orderTimeout

type orderTimeout struct {
	orderService service.IOrderService
	redis        *redis.Client
}

func (mq *orderTimeout) Send(orderId string) {
	if err := mq.redis.ZAdd(ctx, orderTimeoutDelayQueue, &redis.Z{
		Score:  float64(time.Now().Unix() + conf.Conf.Order.Expire),
		Member: orderId,
	}).Err(); err != nil {
		log.Printf("订单【%s】加入延迟队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%s】加入延迟队列", orderId)
	}
	return
}

func (mq *orderTimeout) Remove(orderId string) {
	if err := mq.redis.ZRem(ctx, orderTimeoutDelayQueue, orderId).Err(); err != nil {
		log.Printf("订单【%s】移除延迟队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%s】移除延迟队列", orderId)
	}
	return
}

func (mq *orderTimeout) Receive() {
	var (
		list      []string
		err       error
		orderInfo model.OrderInfo
	)
	for {
		if list, err = mq.redis.ZRangeByScore(ctx, orderTimeoutDelayQueue, &redis.ZRangeBy{
			Min:    "0",
			Max:    strconv.FormatInt(time.Now().Unix(), 10),
			Offset: 0,
			Count:  1,
		}).Result(); err != nil {
			log.Printf("redis.ZRangeByScore() failed, err: %v", err)
			continue
		}
		if len(list) == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if orderInfo, err = mq.orderService.GetOrderInfo(list[0]); err != nil {
			log.Printf("orderService.GetOrderInfo() failed, orderId: %s, err: %v", list[0], err)
			return
		}
		if err = mq.orderService.CloseOrder(orderInfo.UserId, orderInfo.OrderId); err != nil {
			log.Printf("orderService.CloseOrder() failed, orderId: %s, err: %v", err)
			return
		}
		log.Printf("订单【%s】已关闭", orderInfo.OrderId)
	}
}
