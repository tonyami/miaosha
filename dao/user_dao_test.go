package dao

import (
	"miaosha/conf"
	"testing"
)

var userDao *UserDao

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	userDao = NewUserDao()
}

func TestUserDao_QueryById(t *testing.T) {
	user, err := userDao.QueryById(5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
}
