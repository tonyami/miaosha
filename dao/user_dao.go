package dao

import (
	"database/sql"
	"miaosha/conf"
	"miaosha/lib/database/mysql"
	"miaosha/model"
)

type IUserDao interface {
	QueryById(id int64) (*model.User, error)
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
	_queryByIdSql = "select id,username,password,status,create_time,update_time from user where id=?"
)

func (d *UserDao) QueryById(id int64) (user *model.User, err error) {
	user = &model.User{}
	if err = d.db.QueryRow(_queryByIdSql, id).Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime, &user.UpdateTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}
