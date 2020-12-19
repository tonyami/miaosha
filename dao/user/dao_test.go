package user

import (
	"log"
	"miaosha/conf"
	"miaosha/model"
	"testing"
	"time"
)

var d *Dao

func init() {
	conf.Init()
	d = New()
}

func TestDao_QueryById(t *testing.T) {
	user, err := d.QueryByMobile("18812345678")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
}

func TestDao_Save(t *testing.T) {
	user := &model.User{
		Mobile:       "18800000000",
		Password:     "123456",
		Salt:         "miaosha",
		RegisterTime: time.Now(),
	}
	id, err := d.Save(user)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("%d", id)
}
