package service

import (
	"log"
	"miaosha/dao"
	"miaosha/lib/code"
	"miaosha/model"
)

type IUserService interface {
	Register(user *model.User) error
}

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}

func (s *UserService) Register(user *model.User) (err error) {
	var (
		u  *model.User
		id int64
	)
	// 根据手机号码查询用户是否已经注册
	if u, err = s.userDao.QueryByMobile(user.Mobile); err != nil {
		log.Printf("【注册】: %s, mobile: %s\n", err, user.Mobile)
		err = code.SystemErr
		return
	}
	if u.Id > 0 {
		err = code.UserMobileRegistered
		return
	}
	u.New(user)
	if id, err = s.userDao.Save(u); err != nil {
		log.Printf("【注册】: %s; User: %#v\n", err, user)
		err = code.SystemErr
		return
	}
	if id == 0 {
		err = code.UserRegisterErr
	}
	return
}
