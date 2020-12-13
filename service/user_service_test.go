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

func TestUserService_GetInfo(t *testing.T) {
	user, err := userService.GetInfo(5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
}
