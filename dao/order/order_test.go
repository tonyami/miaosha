package order

import (
	"miaosha/conf"
	"miaosha/model"
	"miaosha/util/key"
	"testing"
	"time"
)

var d *Dao

func init() {
	conf.Init()
	d = New()
}

func TestDao_Miaosha(t *testing.T) {
	order := &model.Order{
		Id:         key.CreateOrderId(),
		UserId:     1,
		GoodsId:    4,
		GoodsName:  "棒棒糖（荔枝味）",
		GoodsImg:   "/static/sugar.jpg",
		Price:      10,
		CreateTime: time.Now(),
		Status:     conf.OrderPayWaiting,
	}
	if err := d.Miaosha(order); err != nil {
		t.Fatal(err)
	}
}
