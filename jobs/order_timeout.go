package jobs

import (
	"log"
	"miaosha/conf"
	"miaosha/internal/rdb"
	"miaosha/service"
	"strconv"
	"sync"
	"time"
)

const order_timeout_delay_queue = "order_timeout_delay_queue"

var orderTimeoutJobInstance *OrderTimeoutJob

var once sync.Once

func init() {
	log.Printf("init order timeout job ...")
	once.Do(func() {
		orderTimeoutJobInstance = new(OrderTimeoutJob)
	})
}

type OrderTimeoutJob struct {
}

func GetOrderTimeoutJob() *OrderTimeoutJob {
	return orderTimeoutJobInstance
}

func (*OrderTimeoutJob) Add(orderId int64) {
	if err := rdb.ZAdd(order_timeout_delay_queue, float64(time.Now().Unix()+conf.Conf.Order.Expire), orderId); err != nil {
		log.Printf("订单【%d】进入延时队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%d】进入延时队列", orderId)
	}
	return
}

func (*OrderTimeoutJob) Remove(orderId int64) {
	if err := rdb.ZRem(order_timeout_delay_queue, orderId); err != nil {
		log.Printf("订单【%d】移除延时队列失败, err: %v", orderId, err)
	} else {
		log.Printf("订单【%d】移除延时队列", orderId)
	}
	return
}

func (*OrderTimeoutJob) Start() {
	for {
		time.Sleep(500 * time.Millisecond)
		list, err := rdb.ZRangeByScore(order_timeout_delay_queue, "0", strconv.FormatInt(time.Now().Unix(), 10), 0, 0)
		if err != nil {
			continue
		}
		if len(list) > 0 {
			orderId, _ := strconv.ParseInt(list[0], 10, 64)
			if err = service.GetOrderService().SysCancel(orderId); err != nil {
				log.Printf("订单【%d】取消失败, err: %v", orderId, err)
				continue
			}
			log.Printf("订单【%d】已关闭", orderId)
		}
	}
}
