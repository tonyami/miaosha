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

func TestDao_QueryAll(t *testing.T) {
	goodsList, err := d.QueryAll(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, goods := range goodsList {
		t.Logf("%#v", goods)
	}
}

func TestDao_QueryById(t *testing.T) {
	goods, err := d.QueryById(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", goods)
}
