package user

import (
	"errors"
	"log"
	"miaosha/internal/db"
	"miaosha/internal/key"
	"miaosha/service"
	"sync"
)

var once sync.Once

func init() {
	log.Printf("init user service...")
	once.Do(func() {
		service.IUserService = new(userService)
	})
}

type userService struct {
}

func (*userService) SendSmsCode(mobile string) (code string, err error) {
	dao := new(Dao)
	code = key.SmsCode()
	if err = dao.SaveSmsCode(mobile, code); err != nil {
		err = errors.New("验证码发送失败")
	}
	return
}

func (*userService) Login(mobile, smsCode string) (token string, err error) {
	dao := NewDao(db.Get())
	if rdbCode, err := dao.GetSmsCode(mobile); err != nil || rdbCode != smsCode {
		return "", errors.New("验证码错误")
	}
	var u *User
	if u, err = dao.GetByMobile(mobile); err != nil {
		return "", errors.New("db error")
	}
	if u.Id == 0 {
		u = &User{
			Mobile: mobile,
		}
		if u.Id, err = dao.Insert(u); err != nil || u.Id == 0 {
			return "", errors.New("db error")
		}
	}
	token = key.Token()
	if err = dao.SaveToken(token, u); err != nil {
		return "", errors.New("redis error")
	}
	if err = dao.DeleteSmsCode(mobile); err != nil {
		return "", errors.New("redis error")
	}
	return
}

func (*userService) GetUserByToken(token string) (user *service.UserDTO, err error) {
	dao := new(Dao)
	u, err := dao.GetToken(token)
	if err != nil {
		return
	}
	user = u.toDTO()
	return
}

func (*userService) RenewToken(token string) (err error) {
	dao := new(Dao)
	err = dao.RenewToken(token)
	return
}
