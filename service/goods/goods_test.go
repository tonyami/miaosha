package goods

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"testing"
)

var s *Service

func init() {
	conf.Init()
	db.Init()
	s = New()
}

func TestService_GetList(t *testing.T) {
	if goods, err := s.GetList(1); err != nil || len(goods) == 0 {
		t.Fatal(err)
	}
}

func TestService_Get(t *testing.T) {
	if goods, err := s.Get(1); err != nil || goods == nil {
		t.Fatal(err)
	}
}
