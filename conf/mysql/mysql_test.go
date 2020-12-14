package mysql

import (
	"testing"
)

func TestNew(t *testing.T) {
	db1 := New()
	db2 := New()
	t.Log(db1 == db2)
}
