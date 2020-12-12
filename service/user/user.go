package user

import (
	"log"
	"miaosha/lib/code"
	"miaosha/model"
)

func (s *Service) GetInfo(id int64) (user *model.User, err error) {
	if user, err = s.d.QueryById(id); err != nil {
		log.Println(err)
		err = code.SystemErr
		return
	}
	if user.Id == 0 {
		err = code.UserNotFound
	}
	return
}
