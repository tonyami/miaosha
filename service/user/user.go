package user

import (
	"encoding/json"
	"fmt"
	"log"
	"miaosha/conf"
	"miaosha/dao/user"
	"miaosha/internal/cache"
	"miaosha/internal/code"
	"miaosha/internal/key"
	"miaosha/model"
	"time"
)

type Service struct {
	dao *user.Dao
}

func New() *Service {
	return &Service{
		dao: user.New(),
	}
}

func (s *Service) SendSmsCode(mobile string) (smsCode string, err error) {
	// 1、生成验证码
	smsCode = key.SmsCode()
	// 2、保存验证码
	if err = cache.Conn().Set(fmt.Sprintf(conf.SmsCodeKey, mobile), smsCode, conf.SmsCodeIn); err != nil {
		return
	}
	return
}

func (s *Service) Login(mobile, smsCode string) (token string, err error) {
	// 1、尝试从缓存中获取验证码
	var c string
	if c, err = cache.Conn().Get(fmt.Sprintf(conf.SmsCodeKey, mobile)); err != nil {
		return
	}
	// 2、比较缓存中验证码和输入的验证码
	if c != smsCode {
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
			CreateTime: time.Now(),
		}
		if u.Id, err = s.dao.Insert(u); err != nil {
			log.Printf("Login Failed: %s", err)
			err = code.SystemErr
			return
		}
	}
	// 5、生成token
	token = key.Token()
	// 6、序列化用户对象，并保存到缓存中
	bytes, _ := json.Marshal(u)
	if err = cache.Conn().Set(fmt.Sprintf(conf.TokenKey, token), string(bytes), conf.TokenIn); err != nil {
		return
	}
	// 7、使缓存中验证码立即失效
	if err = cache.Conn().Expire(fmt.Sprintf(conf.SmsCodeKey, u.Mobile), 0); err != nil {
		return
	}
	return
}

func (s *Service) Auth(token string) (user *model.User, err error) {
	var val string
	if val, err = cache.Conn().Get(fmt.Sprintf(conf.TokenKey, token)); err != nil {
		return
	}
	if len(val) == 0 {
		err = code.Unauthorized
		return
	}
	if err = json.Unmarshal([]byte(val), &user); err != nil {
		err = code.SystemErr
	}
	// 重置 token 有效期
	if err = cache.Conn().Expire(fmt.Sprintf(conf.TokenKey, token), conf.TokenIn); err != nil {
		return
	}
	return
}
