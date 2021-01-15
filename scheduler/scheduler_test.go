package scheduler

import (
	"miaosha/conf"
	"miaosha/internal/db"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	conf.Init()
	db.Init()
	Init()
	time.Sleep(10 * time.Second)
}
