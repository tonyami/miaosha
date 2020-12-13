package model

import (
	"miaosha/lib/code"
	"miaosha/lib/encrypt"
	"miaosha/lib/random"
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	Mobile       string    `json:"mobile"`
	Password     string    `json:"-"`
	Salt         string    `json:"-"`
	RegisterTime time.Time `json:"register_time"`
}

func (user *User) Check() (err error) {
	// 校验手机长度是否等于11位，密码是否小于6位
	if user.Mobile == "" || len(user.Mobile) != 11 || user.Password == "" || len(user.Password) < 6 {
		err = code.InvalidParams
	}
	return
}

func (user *User) New(u *User) {
	user.Salt = createSalt()
	user.Mobile = u.Mobile
	user.Password = encrypt.Md5(u.Password + user.Salt)
	user.RegisterTime = time.Now()
	return
}

func createSalt() string {
	return random.Create(16)
}
