package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"miaosha/internal/rdb"
)

const (
	smsCodeKey    = "sms_code_%s"
	smsCodeExpire = 30
	tokenKey      = "token_%s"
	tokenExpire   = 3600
)

func (*Dao) SaveSmsCode(mobile, smsCode string) error {
	return rdb.Set(fmt.Sprintf(smsCodeKey, mobile), smsCode, smsCodeExpire)
}

func (*Dao) GetSmsCode(mobile string) (string, error) {
	return rdb.Get(fmt.Sprintf(smsCodeKey, mobile))
}

func (*Dao) DeleteSmsCode(mobile string) error {
	return rdb.Expire(fmt.Sprintf(smsCodeKey, mobile), 0)
}

func (*Dao) SaveToken(token string, user *User) (err error) {
	data, err := json.Marshal(user)
	if err != nil {
		log.Printf("json.Marshal(%#v) failed, err: %v", user, err)
		return
	}
	return rdb.Set(fmt.Sprintf(tokenKey, token), string(data), tokenExpire)
}

func (*Dao) GetToken(token string) (user *User, err error) {
	data, err := rdb.Get(fmt.Sprintf(tokenKey, token))
	if err != nil {
		return
	}
	if len(data) == 0 {
		return nil, errors.New("token error")
	}
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		log.Printf("json.Unmarshal(%s) failed, err: %v", data, err)
		return
	}
	return
}

func (dao *Dao) RenewToken(token string) (err error) {
	return rdb.Expire(fmt.Sprintf(tokenKey, token), tokenExpire)
}
