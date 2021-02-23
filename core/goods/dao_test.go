package goods

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

func TestDao_GetList(t *testing.T) {
	if list, err := dao.GetList(1, 10); err != nil || len(list) == 0 {
		t.Fatal(err)
	}
}

func TestDao_Get(t *testing.T) {
	if goods, err := dao.Get(1); err != nil || goods.Id == 0 {
		t.Fatal(err)
	}
}
