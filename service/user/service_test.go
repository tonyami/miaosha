package user

import (
	"miaosha/model"
	"testing"
)

var s = New()

func TestService_Register(t *testing.T) {
	user := &model.User{
		Mobile:   "18800001111",
		Password: "123456",
	}
	if err := s.Register(user); err != nil {
		t.Fatal(err)
	}
}

func TestService_Login(t *testing.T) {
	user := &model.User{
		Mobile:   "18800001111",
		Password: "123456",
	}
	token, err := s.Login(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", token)
}
