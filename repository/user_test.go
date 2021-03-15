package repository

import (
	"miaosha/conf"
	"miaosha/infra/db"
	"miaosha/infra/rdb"
	"miaosha/model"
	"testing"
)

func init() {
	if err := conf.Init(""); err != nil {
		panic(err)
	}
	if err := db.Init(); err != nil {
		panic(err)
	}
	if err := rdb.Init(); err != nil {
		panic(err)
	}
}

func TestSaveLoginSmsCode(t *testing.T) {
	if err := SaveLoginSmsCode("18812345678", "123456"); err != nil {
		t.Fatal(err)
	}
}

func TestGetLoginSmsCode(t *testing.T) {
	code, err := GetLoginSmsCode("18812345678")
	if err != nil {
		t.Fatal(err)
	}
	if code != "123456" {
		t.Fatalf("验证码错误, code: %s", code)
	}
}

func TestDeleteLoginSmsCode(t *testing.T) {
	if err := DeleteLoginSmsCode("188123456"); err != nil {
		t.Fatal(err)
	}
}

func TestGetUserByMobile(t *testing.T) {
	if user, err := GetUserByMobile("18800000001"); err != nil || user.Id == 0 {
		t.Fatal(err)
	}
}

func TestSaveUser(t *testing.T) {
	user := model.User{
		Mobile: "18800000009",
	}
	if id, err := SaveUser(user); err != nil || id == 0 {
		t.Fatal(err)
	}
}
