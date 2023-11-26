package mq

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/service"
	"time"
)

const precreateOrderKey = "precreate_order_queue"

var PrecreateOrder *precreateOrder

type precreateOrder struct {
	orderService service.IOrderService
	redis        *redis.Client
}

type PrecreateOrderMsg struct {
	UserId  int64
	GoodsId int64
}

func (mq *precreateOrder) Send(msg PrecreateOrderMsg) error {
	data, _ := json.Marshal(msg)
	return mq.redis.LPush(ctx, precreateOrderKey, data).Err()
}

func (mq *precreateOrder) Receive() {
	var (
		popStr string
		err    error
	)
	for {
		if popStr, err = mq.redis.RPop(ctx, precreateOrderKey).Result(); err != nil {
			if err == redis.Nil {
				err = nil
			} else {
				log.Printf("redis.BRPop() failed, err: %v", err)
				return
			}
		}
		if len(popStr) == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		var msg PrecreateOrderMsg
		if err = json.Unmarshal([]byte(popStr), &msg); err != nil {
			continue
		}
		if msg.GoodsId == 0 || msg.UserId == 0 {
			continue
		}
		if err = mq.orderService.CreateOrder(msg.UserId, msg.GoodsId); err != nil {
			log.Printf("orderService.CreateOrder() failed, err: %v", err)
			return
		}
		if err = mq.orderService.UnLock(msg.UserId, msg.GoodsId); err != nil {
			log.Printf("orderService.UnLocks() failed, err: %v", err)
			return
		}
	}
}
