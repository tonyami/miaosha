package model

import (
	"miaosha/conf/errmsg"
	"miaosha/util"
	"time"
)

type User struct {
	Id           int64
	Mobile       string
	Password     string
	Salt         string
	RegisterTime time.Time
}

type UserDTO struct {
	Id           int64  `json:"id"`
	Mobile       string `json:"mobile"`
	RegisterTime string `json:"registerTime"`
}

func (user *User) ToDTO() (dto *UserDTO) {
	dto = &UserDTO{}
	dto.Id = user.Id
	dto.Mobile = user.Mobile
	dto.RegisterTime = user.RegisterTime.Format("2006-01-02 15:04:05")
	return
}

func (user *User) Check() (err error) {
	// 校验手机长度是否等于11位，密码是否小于6位
	if user.Mobile == "" || len(user.Mobile) != 11 || user.Password == "" || len(user.Password) < 6 {
		err = errmsg.InvalidParameter
	}
	return
}

func (user *User) EqualsPwd(u *User) (err error) {
	if user.Password != util.Md5(u.Password+user.Salt) {
		return errmsg.PasswordErr
	}
	return
}

func (user *User) New(u *User) {
	user.Salt = util.CreateSalt()
	user.Mobile = u.Mobile
	user.Password = util.Md5(u.Password + user.Salt)
	user.RegisterTime = time.Now()
	return
}
