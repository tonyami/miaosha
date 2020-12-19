package mysql

import (
	"database/sql"
	"fmt"
	"miaosha/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func New() *sql.DB {
	once.Do(func() {
		db = connect()
	})
	return db
}

func connect() (db *sql.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/miaosha?charset=utf8&parseTime=true&loc=Local", conf.Conf.DbUser, conf.Conf.DbPassword, conf.Conf.DbAddr)
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
	return
}
