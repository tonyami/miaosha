package scheduler

import (
	"log"
	"miaosha/conf"
	"miaosha/dao/order"
	"time"
)

var (
	orderChannelSize = 10 // 超时订单缓冲区大小
	orderConsumerNum = 5  // 消费者数量
	orderDao         *order.Dao
)

// 生产者：查询超时订单
func orderProducer(orderChannel chan<- string) {
	for {
		ids, err := orderDao.GetOvertimeList(conf.OrderExpire)
		if err != nil {
			log.Printf("【Scheduler】GetOvertimeList Failed: %s", err)
		}
		for i := range ids {
			log.Printf("【Scheduler】订单超时: %s", ids[i])
			orderChannel <- ids[i]
		}
		// 指定时间扫描一次
		time.Sleep(conf.OrderSchedulerInterval * time.Second)
	}
}

// 消费者：修改订单状态、加库存
func orderConsumer(orderChannel <-chan string) {
	for id := range orderChannel {
		o, err := orderDao.Get(id)
		if err != nil {
			log.Printf("【Scheduler】订单超时，处理失败: %s", err)
		}
		if err = orderDao.Close(o); err != nil {
			log.Printf("【Scheduler】订单超时，处理失败: %s", err)
		}
		log.Printf("【Scheduler】订单超时，处理成功: %s", id)
	}
}

func handleOrderOvertime() {
	orderDao = order.New()
	// 待处理超时订单
	var orderChannel = make(chan string, orderChannelSize)
	// 创建生产者
	go orderProducer(orderChannel)
	// 创建一组消费者
	for i := 0; i < orderConsumerNum; i++ {
		go orderConsumer(orderChannel)
	}
}
