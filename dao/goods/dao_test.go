package goods

import (
	"miaosha/conf"
	"testing"
)

var d *Dao

func init() {
	conf.Init()
	d = New()
}

func TestDao_GetList(t *testing.T) {
	list, err := d.GetList(1, 10)
	if err != nil || len(list) == 0 {
		t.Fatal()
	}
}

func TestDao_Get(t *testing.T) {
	goods, err := d.Get(1)
	if err != nil || goods.Goods.Id == 0 {
		t.Fatal()
	}
}
