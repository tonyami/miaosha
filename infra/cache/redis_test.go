package cache

import (
	"miaosha/conf"
	"testing"
)

func init() {
	conf.Init("")
}

func TestInit(t *testing.T) {
	Init()
}
