package jobs

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

const order_timeout_delay_queue = "order_timeout_delay_queue"

var (
	orderTimeJob *OrderTimeoutJob
	once         sync.Once
	ctx          = context.Background()
)

func init() {
	once.Do(func() {
		log.Print("init order timeout job...")
		orderTimeJob = new(OrderTimeoutJob)
	})
}

type OrderTimeoutJob struct {
}

func GetOrderTimeoutJob() *OrderTimeoutJob {
	return orderTimeJob
}

func (*OrderTimeoutJob) Add(orderId string) {
	if err := rdb.Conn().ZAdd(ctx, order_timeout_delay_queue, &redis.Z{
		Score:  float64(time.Now().Unix() + conf.Conf.Order.Expire),
		Member: orderId,
	}).Err(); err != nil {
		log.Printf("订单【%s】进入延时队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%s】进入延时队列", orderId)
	}
	return
}

func (*OrderTimeoutJob) Remove(orderId string) {
	if err := rdb.Conn().ZRem(ctx, order_timeout_delay_queue, orderId).Err(); err != nil {
		log.Printf("订单【%s】移除延时队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%s】移除延时队列", orderId)
	}
	return
}

func (*OrderTimeoutJob) Start() {
	for {
		list, err := rdb.Conn().ZRangeByScore(ctx, order_timeout_delay_queue, &redis.ZRangeBy{
			Min:    "0",
			Max:    strconv.FormatInt(time.Now().Unix(), 10),
			Offset: 0,
			Count:  1,
		}).Result()
		if err != nil {
			continue
		}
		if len(list) > 0 {
			order, err := repository.GetOrderByOrderId(list[0])
			if err != nil {
				log.Printf("GetOrderByOrderId(%s) failed, err: %v", list[0], err)
				continue
			}
			if err = repository.CloseOrder(order.OrderId, order.GoodsId); err != nil {
				log.Printf("CloseOrder(%s, %d) failed, err: %v", order.OrderId, order.GoodsId, err)
				continue
			}
			GetOrderTimeoutJob().Remove(order.OrderId)
			log.Printf("订单【%s】已关闭", order.OrderId)
		}
	}
}
