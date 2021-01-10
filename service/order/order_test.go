package order

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"miaosha/service/goods"
	"testing"
)

var s *Service

func init() {
	conf.Init()
	db.Init()
	goodsService := goods.New()
	s = New(goodsService)
}

func TestService_Get(t *testing.T) {
	if order, err := s.Get("2101086726742339", 1); err != nil || order == nil {
		t.Fatal(err)
	}
}

func TestService_GetList(t *testing.T) {
	if orders, err := s.GetList(1, 1, "0"); err != nil || len(orders) == 0 {
		t.Fatal(err)
	}
}

func TestService_Miaosha(t *testing.T) {
	if _, err := s.Miaosha(1, 6); err != nil {
		t.Fatal(err)
	}
}
