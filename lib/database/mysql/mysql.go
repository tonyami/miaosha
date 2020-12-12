package mysql

import (
	"database/sql"
	"miaosha/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	instance *sql.DB
	once     sync.Once
)

func New(c *conf.Mysql) *sql.DB {
	once.Do(func() {
		instance = connect(c)
	})
	return instance
}

func connect(c *conf.Mysql) (db *sql.DB) {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(c.Idles)
	db.SetMaxOpenConns(c.Opens)
	db.SetConnMaxLifetime(c.Lifetime)
	return
}
