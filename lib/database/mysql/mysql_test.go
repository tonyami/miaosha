package mysql

import (
	"miaosha/conf"
	"testing"
)

var c *conf.MysqlConf

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	c = conf.Mysql
}

func TestNew(t *testing.T) {
	db1 := New(c)
	db2 := New(c)
	t.Log(db1 == db2)
}
