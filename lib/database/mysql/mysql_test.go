package mysql

import (
	"miaosha/conf"
	"testing"
)

var c *conf.Config

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	c = conf.Conf
}

func TestNew(t *testing.T) {
	db1 := New(c.Mysql)
	db2 := New(c.Mysql)
	t.Log(db1 == db2)
}
