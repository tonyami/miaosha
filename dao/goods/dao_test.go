package goods

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"testing"
)

var d *Dao

func init() {
	conf.Init()
	db.Init()
	d = New()
}

func TestDao_GetList(t *testing.T) {
	if _, err := d.GetList(1, 10); err != nil {
		t.Fatal(err)
	}
}

func TestDao_Get(t *testing.T) {
	goods, err := d.Get(1)
	if err != nil || goods == nil {
		t.Fatal(err)
	}
}
