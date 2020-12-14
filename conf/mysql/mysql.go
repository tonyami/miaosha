package mysql

import (
	"database/sql"
	"miaosha/conf"
	"os"
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
	dsn := os.Getenv(conf.ENV_MIAOSHA_DSN)
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
