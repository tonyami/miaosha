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

func New(c *conf.MysqlConf) *sql.DB {
	once.Do(func() {
		instance = connect(c)
	})
	return instance
}

func connect(c *conf.MysqlConf) (db *sql.DB) {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	if c.Idles == 0 {
		c.Idles = 2
	}
	db.SetMaxIdleConns(c.Idles)
	if c.Opens == 0 {
		c.Opens = 5
	}
	db.SetMaxOpenConns(c.Opens)
	if c.Lifetime == 0 {
		c.Lifetime = 5000
	}
	db.SetConnMaxLifetime(c.Lifetime)
	return
}
