package user

import (
	"miaosha/service"
	"time"
)

type User struct {
	Id         int64
	Mobile     string
	CreateTime time.Time
	UpdateTime time.Time
}

func (u *User) toDTO() *service.UserDTO {
	return &service.UserDTO{
		Id:           u.Id,
		Mobile:       u.Mobile,
		RegisterTime: u.CreateTime,
	}
}
