package db

import (
	"miaosha/conf"
	"testing"
)

func TestInit(t *testing.T) {
	conf.Init()
	Init()
}
