package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miaosha/conf"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func New(c *conf.DB) *sql.DB {
	once.Do(func() {
		db = open(c)
	})
	return db
}

func open(c *conf.DB) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(3000)
	return db
}
