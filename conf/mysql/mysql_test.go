package mysql

import (
	"miaosha/conf"
	"testing"
)

func init() {
	conf.Init()
}

func TestNew(t *testing.T) {
	t.Log(New())
}
