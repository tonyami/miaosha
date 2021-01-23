package order

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"miaosha/internal/key"
	"miaosha/model"
	"testing"
	"time"
)

var d *Dao

func init() {
	conf.Init()
	db.Init()
	d = New()
}

func TestDao_GetOvertimeList(t *testing.T) {
	ids, err := d.GetOvertimeList(60)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", ids)
}

func TestDao_Close(t *testing.T) {
	order := model.Order{
		Id:      "2101082621259341",
		Status:  conf.OrderWaitPay,
		GoodsId: 4,
	}
	if err := d.Close(&order); err != nil {
		t.Fatal(err)
	}
}

func TestDao_Get(t *testing.T) {
	if order, err := d.Get("2101086726742339"); err != nil || order == nil {
		t.Fatal(err)
	}
}

func TestDao_GetDTO(t *testing.T) {
	if order, err := d.GetDTO("2101086726742339"); err != nil || order == nil {
		t.Fatal(err)
	}
}

func TestDao_GetList(t *testing.T) {
	if orders, err := d.GetList(1, 1, 10, ""); err != nil || len(orders) == 0 {
		t.Fatal(err)
	}
}

func TestDao_Count(t *testing.T) {
	if count, err := d.Count(1, 4); err != nil || count == 0 {
		t.Fatal(err)
	}
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
