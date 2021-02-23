package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"miaosha/conf"
	"time"
)

var db *sql.DB

func Init(c *conf.DB) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Name)
	if db, err = sql.Open("mysql", dsn); err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	db.SetMaxIdleConns(c.Idles)
	db.SetMaxOpenConns(c.Opens)
	db.SetConnMaxLifetime(time.Duration(c.LifeTime) * time.Hour)
	return
}

func Get() *sql.DB {
	return db
}

func Tx(conn *sql.DB, f func(*sql.DB) error) (err error) {
	tx, err := conn.Begin()
	if err != nil {
		log.Printf("db.Begin() failed, err: %v", err)
		return
	}
	defer tx.Rollback()
	if err = f(conn); err != nil {
		log.Printf("db.Tx.f(conn) failed, err: %v", err)
		return
	}
	err = tx.Commit()
	return
}
