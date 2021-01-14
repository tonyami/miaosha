package user

import (
	"database/sql"
	"miaosha/internal/db"
	"miaosha/model"
)

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	return &Dao{
		db: db.Conn(),
	}
}

var (
	_getSql    = `select id, mobile, create_time from miaosha_user where mobile = ? limit 1`
	_insertSql = `insert into miaosha_user(mobile, create_time) values(?, ?)`
)

func (d *Dao) Get(mobile string) (user *model.User, err error) {
	user = &model.User{}
	if err = d.db.QueryRow(_getSql, mobile).Scan(&user.Id, &user.Mobile, &user.CreateTime); err != nil {
		user = nil
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return
}

func (d *Dao) Insert(user *model.User) (insertId int64, err error) {
	var (
		stmt *sql.Stmt
		rs   sql.Result
	)
	if stmt, err = d.db.Prepare(_insertSql); err != nil {
		return
	}
	defer stmt.Close()
	if rs, err = stmt.Exec(user.Mobile, user.CreateTime); err != nil {
		return
	}
	if insertId, err = rs.LastInsertId(); err != nil {
		return
	}
	if insertId == 0 {
		err = sql.ErrNoRows
	}
	return
}
