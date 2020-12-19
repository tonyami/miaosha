package user

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"miaosha/conf/errmsg"
	"miaosha/conf/rdb"
	"miaosha/dao/user"
	"miaosha/model"
	"miaosha/util"
	"time"
)

type Service struct {
	dao *user.Dao
	rdb *redis.Client
}

func New() *Service {
	return &Service{
		dao: user.New(),
		rdb: rdb.New(),
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
		err = errmsg.MobileNotRegistered
		return
	}
	// 密码比对
	if err = u.EqualsPwd(user); err != nil {
		return
	}
	// 登录成功，生成token并写入redis
	token = util.Create(64)
	dto := u.ToDTO()
	data, err := json.Marshal(dto)
	if err != nil {
		err = errmsg.SerializeErr
		return
	}
	s.rdb.Set(context.Background(), conf.TokenPrefix+token, string(data), conf.TokenExpire*time.Second)
	return
}

func (s *Service) Info(token string) (dto *model.UserDTO, err error) {
	res := s.rdb.Get(context.Background(), conf.TokenPrefix+token)
	dto = &model.UserDTO{}
	if res.Err() != nil {
		if res.Err() == redis.Nil {
			err = nil
			return
		}
		log.Printf("【Session验证失败】: %s", res.Err())
		err = errmsg.SystemErr
		return
	}
	if err = json.Unmarshal([]byte(res.Val()), dto); err != nil {
		err = errmsg.SerializeErr
	}
	return
}
