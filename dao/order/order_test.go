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
		Id:         key.OrderId(),
		UserId:     1,
		GoodsId:    5,
		CreateTime: time.Now(),
		Status:     conf.OrderUnPaid,
	}
	if err := d.Miaosha(order); err != nil {
		t.Fatal(err)
	}
}
