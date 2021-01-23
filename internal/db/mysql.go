package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miaosha/conf"
)

var db *sql.DB

func Init() {
	var err error
	if db, err = open(conf.Conf.DB); err != nil {
		panic(err)
	}
}

func Conn() *sql.DB {
	return db
}

func open(c *conf.DB) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Name)
	if db, err = sql.Open("mysql", dsn); err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(3000)
	return
}
