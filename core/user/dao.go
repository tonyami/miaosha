package user

import (
	"database/sql"
	"log"
)

type Dao struct {
	db *sql.DB
}

func NewDao(db *sql.DB) *Dao {
	return &Dao{
		db: db,
	}
}

var (
	_getByMobileSql = "select `id`, `mobile`, `create_time` from `miaosha_user` where `mobile` = ? limit 1"
	_insertSql      = "insert into `miaosha_user`(`mobile`) values(?)"
)

func (dao *Dao) GetByMobile(mobile string) (user *User, err error) {
	user = &User{}
	if err = dao.db.QueryRow(_getByMobileSql, mobile).Scan(&user.Id, &user.Mobile, &user.CreateTime); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("dao.db.QueryRow(_getByMobileSql, %s) failed, err: %v", mobile, err)
		}
	}
	return
}

func (dao *Dao) Insert(user *User) (insertId int64, err error) {
	var (
		stmt *sql.Stmt
		rs   sql.Result
	)
	if stmt, err = dao.db.Prepare(_insertSql); err != nil {
		log.Printf("dao.db.Prepare(_insertSql) failed, err: %v", err)
		return
	}
	defer stmt.Close()
	if rs, err = stmt.Exec(user.Mobile); err != nil {
		log.Printf("stmt.Exec(_insertSql, %s) failed, err: %v", user.Mobile, err)
		return
	}
	return rs.LastInsertId()
}
