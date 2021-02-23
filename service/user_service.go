package service

import (
	"time"
)

var IUserService UserService

func GetUserService() UserService {
	return IUserService
}

type UserService interface {
	SendSmsCode(string) (string, error)
	Login(string, string) (string, error)
	GetUserByToken(string) (*UserDTO, error)
	RenewToken(string) error
}

type UserDTO struct {
	Id           int64     `json:"id"`
	Mobile       string    `json:"mobile"`
	RegisterTime time.Time `json:"registerTime"`
}
