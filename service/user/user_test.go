package user

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

func TestService_Login(t *testing.T) {
	if _, err := s.Login("18800000000", "123456"); err != nil {
		t.Fatal(err)
	}
}
