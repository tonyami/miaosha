package user

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"miaosha/model"
	"testing"
	"time"
)

var d *Dao

func init() {
	conf.Init()
	db.Init()
	d = New()
}

func TestDao_Get(t *testing.T) {
	if user, err := d.Get("18800000000"); err != nil || user == nil {
		t.Fatal(err)
	}
}

func TestDao_Insert(t *testing.T) {
	user := &model.User{
		Mobile:     "18800000004",
		CreateTime: time.Now(),
	}
	if insertId, err := d.Insert(user); err != nil || insertId == 0 {
		t.Fatal(err)
	}
}
