package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"miaosha/infra/db"
	"miaosha/model"
)

type UserRepository interface {
	GetUserByMobile(mobile string) (user model.User, err error)
	SaveUser(user model.User) (id int64, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: db.DB}
}

func (r *userRepository) GetUserByMobile(mobile string) (user model.User, err error) {
	if err = r.db.Get(&user, "select * from miaosha_user where mobile = ?", mobile); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("db.Get() failed, err: %v", err)
		}
	}
	return
}

func (r *userRepository) SaveUser(user model.User) (id int64, err error) {
	var result sql.Result
	if result, err = r.db.Exec("insert into miaosha_user(mobile) values(?)", user.Mobile); err != nil {
		log.Printf("db.Exec() failed, err: %v", err)
		return
	}
	return result.LastInsertId()
}
