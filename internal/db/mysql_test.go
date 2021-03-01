package db

import (
	"database/sql"
	"miaosha/conf"
	"testing"
)

func init() {
	conf.Init()
	if err := Init(conf.Conf.DB); err != nil {
		panic(err)
	}
}

func TestTx(t *testing.T) {
	if err := Tx(Get(), func(conn *sql.DB) (err error) {
		if _, err = conn.Exec("update `tx_user` set `money` = `money` - ? where `id` = ? and `money` >= 0", 100, 1); err != nil {
			return
		}
		if _, err = conn.Exec("update `tx_user` set `money` = `money` + ? where `id` = ? and `money` >= 0", 100, 2); err != nil {
			return
		}
		return
	}); err != nil {
		t.Fatal(err)
	}
}
