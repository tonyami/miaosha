package dao

import (
	"database/sql"
	"miaosha/conf/mysql"
	"miaosha/model"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{
		db: mysql.New(),
	}
}

var (
	_queryByMobileSql = "select id, mobile, password, salt, register_time from user where mobile = ?"
	_saveSql          = "insert into user(mobile, password, salt, register_time) values(?, ?, ?, ?)"
)

func (d *Dao) QueryByMobile(mobile string) (user *model.User, err error) {
	user = &model.User{}
	if err = d.db.QueryRow(_queryByMobileSql, mobile).Scan(&user.Id, &user.Mobile, &user.Password, &user.Salt, &user.RegisterTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}

func (d *Dao) Save(user *model.User) (id int64, err error) {
	stmt, err := d.db.Prepare(_saveSql)
	if err != nil {
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec(user.Mobile, user.Password, user.Salt, user.RegisterTime)
	if err != nil {
		return
	}
	return ret.LastInsertId()
}
