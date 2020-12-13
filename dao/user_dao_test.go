package dao

import (
	"log"
	"miaosha/conf"
	"miaosha/model"
	"testing"
	"time"
)

var userDao *UserDao

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	userDao = NewUserDao()
}

func TestUserDao_QueryById(t *testing.T) {
	user, err := userDao.QueryByMobile("18812345678")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
}

func TestUserDao_Save(t *testing.T) {
	user := &model.User{
		Mobile:       "18800000000",
		Password:     "123456",
		Salt:         "miaosha",
		RegisterTime: time.Now(),
	}
	id, err := userDao.Save(user)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("%d", id)
}
