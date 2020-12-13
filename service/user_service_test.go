package service

import (
	"miaosha/conf"
	"testing"
)

var userService *UserService

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	userService = NewUserService()
}

func TestUserService_Register(t *testing.T) {
	if err := userService.Register("18800001111", "123456"); err != nil {
		t.Fatal(err)
	}
}
