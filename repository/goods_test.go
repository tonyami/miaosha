package repository

import "testing"

func TestGetGoodsList(t *testing.T) {
	if list, err := GetGoodsList(1); err != nil || len(list) == 0 {
		t.Fatal(err)
	}
}

func TestGetGoods(t *testing.T) {
	if goods, err := GetGoods(1); err != nil || goods.Id != 1 {
		t.Fatal(err)
	}
}
