package service

import (
	"log"
	"miaosha/dao"
	"miaosha/lib/code"
	"miaosha/model"
)

type IUserService interface {
	GetInfo(id int64) (*model.User, error)
}

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}

func (s *UserService) GetInfo(id int64) (user *model.User, err error) {
	if user, err = s.userDao.QueryById(id); err != nil {
		log.Println(err)
		err = code.SystemErr
		return
	}
	if user.Id == 0 {
		err = code.UserNotFound
	}
	return
}
