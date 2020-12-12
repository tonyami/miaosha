package user

import (
	"database/sql"
	"miaosha/model"
)

var (
	_queryByIdSql = "select id,username,password,status,create_time,update_time from user where id=?"
)

func (d *Dao) QueryById(id int64) (user *model.User, err error) {
	user = &model.User{}
	if err = d.db.QueryRow(_queryByIdSql, id).Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime, &user.UpdateTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}
