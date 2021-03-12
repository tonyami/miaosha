package db

import (
	"miaosha/conf"
	"testing"
)

func init() {
	if err := conf.Init(""); err != nil {
		panic(err)
	}
}

func TestInit(t *testing.T) {
	if err := Init(); err != nil {
		t.Fatal(err)
	}
}
