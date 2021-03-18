package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/infra/cache"
	"miaosha/infra/code"
	"miaosha/infra/util"
	"miaosha/model"
	"miaosha/repository"
	"miaosha/service"
	"sync"
	"time"
)

var once sync.Once

func InitService() {
	once.Do(func() {
		service.UserService = &userService{
			userRepository: repository.NewUserRepository(),
			orderService:   service.GetOrderService(),
			redis:          cache.Client,
		}
	})
}

type userService struct {
	userRepository repository.UserRepository
	orderService   service.IOrderService
	redis          *redis.Client
}

func (s *userService) SaveLoginSmsCode(mobile, smsCode string) (err error) {
	if err = s.redis.Set(service.Ctx, fmt.Sprintf(service.LoginSmsCodeKey, mobile), smsCode, time.Duration(service.LoginSmsCodeExpire)*time.Second).Err(); err != nil {
		log.Printf("redis.Set() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}

func (s *userService) GetLoginSmsCode(mobile string) (smsCode string, err error) {
	if smsCode, err = s.redis.Get(service.Ctx, fmt.Sprintf(service.LoginSmsCodeKey, mobile)).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("redis.Get() failed, err: %v", err)
			err = code.RedisErr
		}
	}
	return
}

func (s *userService) DeleteLoginSmsCode(mobile string) (err error) {
	if err = s.redis.Expire(service.Ctx, fmt.Sprintf(service.LoginSmsCodeKey, mobile), 0).Err(); err != nil {
		log.Printf("redis.Expire() failed, err: %v", err)
	}
	return
}

func (s *userService) SaveUserToken(token string, user model.User) (err error) {
	var data []byte
	if data, err = json.Marshal(user); err != nil {
		log.Printf("json.Marshal() failed, err: %v", err)
		err = code.SerializeErr
		return
	}
	if err = s.redis.Set(service.Ctx, fmt.Sprintf(service.UserTokenKey, token), string(data), time.Duration(service.UserTokenExpire)*time.Second).Err(); err != nil {
		log.Printf("redis.Set() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}

func (s *userService) GetUserByToken(token string) (user model.User, err error) {
	var data string
	if data, err = s.redis.Get(service.Ctx, fmt.Sprintf(service.UserTokenKey, token)).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("redis.Get() failed, err: %v", err)
			err = code.RedisErr
			return
		}
	}
	if len(data) == 0 {
		err = errors.New("token is empty")
		return
	}
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		log.Printf("json.Unmarshal() failed, err: %v", err)
		err = code.SerializeErr
	}
	return
}

func (s *userService) RenewUserToken(token string) (err error) {
	if err = s.redis.Expire(service.Ctx, fmt.Sprintf(service.UserTokenKey, token), time.Duration(service.UserTokenExpire)*time.Second).Err(); err != nil {
		log.Printf("redis.Expire() failed, err: %v", err)
		err = code.RedisErr
	}
	return
}

func (s *userService) GetSmsCode(mobile string) (smsCode string, err error) {
	smsCode = util.CreateKey(util.Number, service.SmsCodeSize)
	if err := s.SaveLoginSmsCode(mobile, smsCode); err != nil {
		err = code.GetSmsCodeErr
	}
	return
}

func (s *userService) Login(mobile string, smsCode string) (token string, err error) {
	var realCode string
	if realCode, err = s.GetLoginSmsCode(mobile); err != nil {
		err = code.RedisErr
		return
	}
	if smsCode != realCode {
		err = code.SmsCodeErr
		return
	}
	// 根据手机号码查询用户
	var user model.User
	if user, err = s.userRepository.GetUserByMobile(mobile); err != nil {
		err = code.DBErr
		return
	}
	// 用户不存在则自动注册
	if user.Id == 0 {
		user = model.User{
			Mobile: mobile,
		}
		if user.Id, err = s.userRepository.SaveUser(user); err != nil || user.Id == 0 {
			err = code.DBErr
			return
		}
	}
	// 生成token并保存
	token = util.CreateKey(util.AlphabetAndNumber, service.UserTokenSize)
	if err = s.SaveUserToken(token, user); err != nil {
		return
	}
	// 删除验证码
	err = s.DeleteLoginSmsCode(mobile)
	return
}

func (s *userService) GetUserInfo(token string) (userInfo model.UserInfoVO, err error) {
	var (
		user  model.User
		count model.OrderCount
	)
	if user, err = s.GetUserByToken(token); err != nil {
		return
	}
	if count, err = s.orderService.CountOrder(user.Id); err != nil {
		return
	}
	userInfo = model.UserInfoVO{
		User:  user,
		Count: count,
	}
	return
}
