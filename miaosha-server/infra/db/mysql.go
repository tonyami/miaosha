package db

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"miaosha/conf"
	"time"
)

var DB *sqlx.DB

func Init() {
	var err error
	c := conf.Conf.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Name)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	if DB, err = sqlx.ConnectContext(ctx, "mysql", dsn); err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	DB.SetMaxIdleConns(c.Idles)
	DB.SetMaxOpenConns(c.Opens)
}
