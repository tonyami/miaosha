package user

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"testing"
)

var dao *Dao

func init() {
	conf.Init()
	if err := db.Init(conf.Conf.DB); err != nil {
		panic(err)
	}
	dao = NewDao(db.Get())
}

func TestDao_GetByMobile(t *testing.T) {
	if user, err := dao.GetByMobile("18800000000"); err != nil || user == nil {
		t.Fatal(err)
	}
}

func TestDao_Insert(t *testing.T) {
	user := &User{
		Mobile: "18800000008",
	}
	if insertId, err := dao.Insert(user); err != nil || insertId == 0 {
		t.Fatal(err)
	}
}
