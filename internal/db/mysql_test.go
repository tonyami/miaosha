package db

import (
	"miaosha/conf"
	"testing"
)

func TestNew(t *testing.T) {
	conf.Init()
	New(conf.Conf.DB)
}
