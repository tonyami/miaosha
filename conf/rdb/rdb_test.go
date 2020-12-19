package rdb

import (
	"context"
	"miaosha/conf"
	"testing"
	"time"
)

func init() {
	conf.Init()
}

func TestNew(t *testing.T) {
	rdb := New()
	rdb.Set(context.Background(), "key", "value", 10*time.Second)
}
