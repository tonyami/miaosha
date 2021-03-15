package mq

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"miaosha/infra/rdb"
	"miaosha/repository"
	"strconv"
	"sync"
	"time"
)

const orderTimeoutDelayQueue = "order_timeout_delay_queue"

var (
	OrderTimeout *OrderTimeoutMQ
	once         sync.Once
	ctx          = context.Background()
)

func init() {
	once.Do(func() {
		OrderTimeout = new(OrderTimeoutMQ)
		OrderPrecreate = new(OrderPrecreateMQ)
	})
}

type OrderTimeoutMQ struct {
}

func (*OrderTimeoutMQ) Add(orderId string) {
	if err := rdb.Conn().ZAdd(ctx, orderTimeoutDelayQueue, &redis.Z{
		Score:  float64(time.Now().Unix() + conf.Conf.Order.Expire),
		Member: orderId,
	}).Err(); err != nil {
		log.Printf("订单【%s】加入延迟队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%s】加入延迟队列", orderId)
	}
	return
}

func (*OrderTimeoutMQ) Remove(orderId string) {
	if err := rdb.Conn().ZRem(ctx, orderTimeoutDelayQueue, orderId).Err(); err != nil {
		log.Printf("订单【%s】移除延迟队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%s】移除延迟队列", orderId)
	}
	return
}

func (*OrderTimeoutMQ) Receive() {
	for {
		list, err := rdb.Conn().ZRangeByScore(ctx, orderTimeoutDelayQueue, &redis.ZRangeBy{
			Min:    "0",
			Max:    strconv.FormatInt(time.Now().Unix(), 10),
			Offset: 0,
			Count:  1,
		}).Result()
		if err != nil {
			log.Printf("rdb.ZRangeByScore() failed, err: %v", err)
			continue
		}
		if len(list) == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		order, err := repository.GetOrderByOrderId(list[0])
		if err != nil {
			log.Printf("GetOrderByOrderId() failed, orderId: %s, err: %v", list[0], err)
			return
		}
		if err = repository.CloseOrder(order); err != nil {
			log.Printf("CloseOrder() failed, orderId: %s, err: %v", order.OrderId, err)
			return
		}
		OrderTimeout.Remove(order.OrderId)
		log.Printf("订单【%s】已关闭", order.OrderId)
	}
}
