package dao

import (
	"database/sql"
	"miaosha/conf"
	"miaosha/lib/database/mysql"
	"miaosha/model"
)

type IUserDao interface {
	QueryByMobile(string) (*model.User, error)
	Save(*model.User) (int64, error)
}

type UserDao struct {
	db *sql.DB
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: mysql.New(conf.Mysql),
	}
}

var (
	_queryByMobileSql = "select id, mobile, password, salt, register_time from user where mobile = ?"
	_saveSql          = "insert into user(mobile, password, salt, register_time) values(?, ?, ?, ?)"
)

func (d *UserDao) QueryByMobile(mobile string) (user *model.User, err error) {
	user = &model.User{}
	if err = d.db.QueryRow(_queryByMobileSql, mobile).Scan(&user.Id, &user.Mobile, &user.Password, &user.Salt, &user.RegisterTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}

func (d *UserDao) Save(user *model.User) (id int64, err error) {
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
