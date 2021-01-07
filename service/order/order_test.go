package order

import (
	"miaosha/conf"
	"miaosha/service/goods"
	"testing"
)

var s *Service

func init() {
	conf.Init()
	goodsService := goods.New()
	s = New(goodsService)
}

func TestService_Miaosha(t *testing.T) {
	if order, err := s.Miaosha(1, 4); err != nil || order == nil {
		t.Fatal(err)
	}
}
