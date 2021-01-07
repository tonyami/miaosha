package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/conf"
	"miaosha/dao/user"
	"miaosha/internal/cache"
	"miaosha/internal/code"
	"miaosha/model"
	"miaosha/util/key"
	"time"
)

type Service struct {
	dao      *user.Dao
	redisCli *redis.Client
}

func New() *Service {
	return &Service{
		dao:      user.New(),
		redisCli: cache.New(conf.Conf.Redis),
	}
}

func (s *Service) SendSmsCode(mobile string) (smsCode string, err error) {
	// 1、生成验证码
	smsCode = key.CreateSmsCode()
	// 2、保存验证码
	if sc := s.redisCli.Set(context.Background(), fmt.Sprintf(conf.SmsCodeKey, mobile), smsCode, conf.SmsCodeIn); sc.Err() != nil && sc.Err() != redis.Nil {
		log.Printf("SendSmsCode Failed: %s", sc.Err())
		err = code.SystemErr
		return
	}
	return
}

func (s *Service) Login(mobile, smsCode string) (token string, err error) {
	// 1、尝试从缓存中获取验证码
	sc := s.redisCli.Get(context.Background(), fmt.Sprintf(conf.SmsCodeKey, mobile))
	if sc.Err() != nil && sc.Err() != redis.Nil {
		log.Printf("Login Failed: %s", sc.Err())
		err = code.SystemErr
		return
	}
	// 2、比较缓存中验证码和输入的验证码
	if sc.Val() != smsCode {
		err = code.CodeErr
		return
	}
	var u *model.User
	// 3、根据手机号码从数据库中查询用户
	if u, err = s.dao.Get(mobile); err != nil {
		log.Printf("Login Failed: %s", err)
		err = code.SystemErr
		return
	}
	// 4、如果用户不存在，则注册
	if u == nil {
		u = &model.User{
			Mobile:     mobile,
			Avatar:     conf.DefaultAvatar,
			CreateTime: time.Now(),
		}
		if u.Id, err = s.dao.Insert(u); err != nil {
			log.Printf("Login Failed: %s", err)
			err = code.SystemErr
			return
		}
	}
	// 5、生成token
	token = key.CreateToken()
	// 6、序列化用户对象，并保存到缓存中
	bytes, _ := json.Marshal(u)
	if sc2 := s.redisCli.Set(context.Background(), fmt.Sprintf(conf.TokenKey, token), string(bytes), conf.TokenIn); sc2.Err() != nil && sc2.Err() != redis.Nil {
		log.Printf("Login Failed: %s", sc2)
		err = code.SystemErr
		return
	}
	// 7、删除验证码缓存，使验证码失效
	if sc3 := s.redisCli.Del(context.Background(), fmt.Sprintf(conf.SmsCodeKey, u.Mobile)); sc3.Err() != nil && sc3.Err() != redis.Nil {
		log.Printf("Login Failed: %s", sc3)
		err = code.SystemErr
	}
	return
}

func (s *Service) Auth(token string) (user *model.User, err error) {
	sc := s.redisCli.Get(context.Background(), fmt.Sprintf(conf.TokenKey, token))
	if sc.Err() != nil && sc.Err() != redis.Nil {
		log.Printf("GetUser Failed: %s", sc)
		err = code.SystemErr
		return
	}
	if len(sc.Val()) == 0 {
		err = code.Unauthorized
		return
	}
	if err = json.Unmarshal([]byte(sc.Val()), &user); err != nil {
		err = code.SystemErr
	}
	// 重置 token 有效期
	if sc2 := s.redisCli.Expire(context.Background(), fmt.Sprintf(conf.TokenKey, token), conf.TokenIn); sc2.Err() != nil && sc2.Err() != redis.Nil {
		log.Printf("Auth Failed: %s", sc2)
		err = code.SystemErr
		return
	}
	return
}
