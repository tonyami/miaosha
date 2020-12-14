package redis

import "testing"

func TestNew(t *testing.T) {
	rdb1 := New()
	rdb2 := New()
	t.Log(rdb1 == rdb2)
}
