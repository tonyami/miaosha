package user

import (
	"log"
	"miaosha/conf/errmsg"
	dao "miaosha/dao/user"
	"miaosha/model"
	"miaosha/util"
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{
		dao: dao.New(),
	}
}

func (s *Service) Register(user *model.User) (err error) {
	var (
		u  *model.User
		id int64
	)
	// 根据手机号码查询用户是否已经注册
	if u, err = s.dao.QueryByMobile(user.Mobile); err != nil {
		log.Printf("【注册失败】: %s, %#v\n", err, user)
		err = errmsg.SystemErr
		return
	}
	if u.Id == 0 {
		err = errmsg.MobileNotRegistered
		return
	}
	u.New(user)
	if id, err = s.dao.Save(u); err != nil {
		log.Printf("【注册失败】: %s; %#v\n", err, user)
		err = errmsg.SystemErr
		return
	}
	if id == 0 {
		err = errmsg.RegisterErr
	}
	return
}

func (s *Service) Login(user *model.User) (token string, err error) {
	var (
		u *model.User
	)
	// 根据手机号码查询用户是否已经注册
	if u, err = s.dao.QueryByMobile(user.Mobile); err != nil {
		log.Printf("【登录失败】: %s, %#v\n", err, user)
		err = errmsg.SystemErr
		return
	}
	if u.Id == 0 {
		err = errmsg.MobileRegistered
		return
	}
	// 密码比对
	if err = u.EqualsPwd(user); err != nil {
		return
	}
	// 登录成功，生成token并写入redis
	token = createToken()
	return
}

func createToken() string {
	return util.Create(64)
}
