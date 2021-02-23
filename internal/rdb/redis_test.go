package rdb

import (
	"miaosha/conf"
	"testing"
)

func init() {
	conf.Init()
	if err := Init(conf.Conf.Redis); err != nil {
		panic(err)
	}
}

func TestSet(t *testing.T) {
	if err := Set("key", "value", 60); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	val, err := Get("key")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
}

func TestExpire(t *testing.T) {
	if err := Expire("key", 60); err != nil {
		t.Fatal(err)
	}
}
