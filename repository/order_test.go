package repository

import (
	"testing"
)

func TestGetOrderList(t *testing.T) {
	if list, err := GetOrderList(1, "", 1); err != nil || len(list) == 0 {
		t.Fatal(err)
	}
}

func TestCountOrder(t *testing.T) {
	if count, err := CountOrder(8); err != nil || count.Closed == 0 {
		t.Fatal(err)
	}
}
