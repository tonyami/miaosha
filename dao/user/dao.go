package user

import (
	"database/sql"
	"miaosha/conf"
	"miaosha/lib/database/mysql"
)

type Dao struct {
	db *sql.DB
}

func New(c *conf.Config) *Dao {
	return &Dao{
		db: mysql.New(c.Mysql),
	}
}
