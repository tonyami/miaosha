package db

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"miaosha/conf"
	"time"
)

var conn *sqlx.DB

func Init() (err error) {
	c := conf.Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Name)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	if conn, err = sqlx.ConnectContext(ctx, "mysql", dsn); err != nil {
		return
	}
	if err = conn.Ping(); err != nil {
		return
	}
	conn.SetMaxIdleConns(c.Idles)
	conn.SetMaxOpenConns(c.Opens)
	conn.SetConnMaxLifetime(time.Duration(c.LifeTime) * time.Second)
	return
}

func Conn() *sqlx.DB {
	return conn
}
