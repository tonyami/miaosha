package cache

import (
	"miaosha/conf"
	"testing"
)

func init() {
	conf.Init("../../conf.ini")
}

func TestInit(t *testing.T) {
	Init()
}
