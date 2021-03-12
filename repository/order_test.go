package repository

import (
	"miaosha/model"
	"testing"
)

func TestGetOrderList(t *testing.T) {
	if list, err := GetOrderList(1, "", 1); err != nil || len(list) == 0 {
		t.Fatal(err)
	}
}

func TestGetOrderById(t *testing.T) {
	if order, err := GetOrderById(9); err != nil || order.Id == 0 {
		t.Fatal(err)
	}
}

func TestCreateOrder(t *testing.T) {
	order := model.Order{}
	if orderId, err := CreateOrder(order); err != nil || orderId == 0 {
		t.Fatal(err)
	}
}

func TestCountRepeatableOrder(t *testing.T) {
	if count, err := CountRepeatableOrder(8, 8); err != nil || count > 0 {
		t.Fatal(err)
	}
}

func TestCountOrder(t *testing.T) {
	if count, err := CountOrder(8); err != nil || count.Closed == 0 {
		t.Fatal(err)
	}
}

func TestCloseOrder(t *testing.T) {
	if err := CloseOrder(9, 1); err != nil {
		t.Fatal(err)
	}
}
