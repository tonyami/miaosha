package user

import (
	"miaosha/conf"
	"testing"
)

var s *Service

func init() {
	conf.Init()
	s = New()
}

func TestService_Login(t *testing.T) {
	if token, err := s.Login("18800000000", "123456"); err != nil || len(token) == 0 {
		t.Fatal(err)
	}
}
