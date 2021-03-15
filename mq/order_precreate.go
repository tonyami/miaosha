package mq

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/infra/rdb"
	"miaosha/model"
	"miaosha/repository"
	"time"
)

const orderPrecreateKey = "order_precreate"

var OrderPrecreate *OrderPrecreateMQ

type OrderPrecreateMQ struct {
}

type OrderMessage struct {
	UserId  int64
	GoodsId int64
}

func (*OrderPrecreateMQ) Send(userId, goodsId int64) error {
	orderMsg := OrderMessage{
		UserId:  userId,
		GoodsId: goodsId,
	}
	data, _ := json.Marshal(orderMsg)
	return rdb.Conn().LPush(ctx, orderPrecreateKey, data).Err()
}

func (*OrderPrecreateMQ) Receive() {
	var (
		popStr  string
		err     error
		orderId string
		goods   model.Goods
	)
	for {
		if popStr, err = rdb.Conn().RPop(ctx, orderPrecreateKey).Result(); err != nil {
			if err == redis.Nil {
				err = nil
			} else {
				log.Printf("rdb.BRPop() failed, err: %v", err)
				return
			}
		}
		if len(popStr) == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		var orderMsg OrderMessage
		if err = json.Unmarshal([]byte(popStr), &orderMsg); err != nil {
			continue
		}
		if orderMsg.GoodsId == 0 || orderMsg.UserId == 0 {
			continue
		}
		// 是否重复秒杀
		if orderId, err = repository.GetOrderIdByUidAndGid(orderMsg.UserId, orderMsg.GoodsId); err != nil {
			log.Printf("db.GetOrderIdByUidAndGid() failed, err: %v", err)
			return
		}
		if len(orderId) > 0 {
			continue
		}
		// 创建订单
		if goods, err = repository.GetGoods(orderMsg.GoodsId); err != nil {
			log.Printf("db.GetGoods() failed, err: %v", err)
			return
		}
		orderInfo := model.NewOrderInfo(orderMsg.UserId, goods)
		if err = repository.CreateOrder(orderInfo); err != nil {
			log.Printf("db.CreateOrder() failed, err: %v", err)
			return
		}
		// 加入延迟队列
		OrderTimeout.Add(orderInfo.OrderId)
	}
}
