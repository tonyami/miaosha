package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"miaosha/infra/db"
	"miaosha/infra/rdb"
	"miaosha/model"
	"time"
)

const (
	loginSmsCodeKey    = "login:sms_code:%s"
	loginSmsCodeExpire = 30
	userTokenKey       = "user:token:%s"
	userTokenExpire    = 3600
)

var ctx = context.Background()

func SaveLoginSmsCode(mobile, code string) (err error) {
	if err = rdb.Conn().Set(ctx, fmt.Sprintf(loginSmsCodeKey, mobile), code, time.Duration(loginSmsCodeExpire)*time.Second).Err(); err != nil {
		log.Printf("dao.SaveLoginSmsCode() failed, err: %v", err)
	}
	return
}

func GetLoginSmsCode(mobile string) (code string, err error) {
	if code, err = rdb.Conn().Get(ctx, fmt.Sprintf(loginSmsCodeKey, mobile)).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("dao.GetLoginSmsCode() failed, err: %v", err)
		}
	}
	return
}

func DeleteLoginSmsCode(mobile string) (err error) {
	if err = rdb.Conn().Expire(ctx, fmt.Sprintf(loginSmsCodeKey, mobile), 0).Err(); err != nil {
		log.Printf("dao.DeleteLoginSmsCode() failed, err: %v", err)
	}
	return
}

func SaveUserToken(token string, user model.User) (err error) {
	var data []byte
	if data, err = json.Marshal(user); err != nil {
		log.Printf("json.Marshal(%#v) failed, err: %v", user, err)
		return
	}
	if err = rdb.Conn().Set(ctx, fmt.Sprintf(userTokenKey, token), string(data), time.Duration(userTokenExpire)*time.Second).Err(); err != nil {
		log.Printf("dao.SaveUserToken() failed, err: %v", err)
	}
	return
}

func GetUserByToken(token string) (user model.User, err error) {
	var data string
	if data, err = rdb.Conn().Get(ctx, fmt.Sprintf(userTokenKey, token)).Result(); err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			log.Printf("dao.SaveUserToken() failed, err: %v", err)
			return
		}
	}
	if len(data) == 0 {
		err = errors.New("token error")
		return
	}
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		log.Printf("json.Unmarshal(%s) failed, err: %v", data, err)
		return
	}
	return
}

func RenewUserToken(token string) (err error) {
	if err = rdb.Conn().Expire(ctx, fmt.Sprintf(userTokenKey, token), time.Duration(userTokenExpire)*time.Second).Err(); err != nil {
		log.Printf("dao.RenewUserToken() failed, err: %v", err)
	}
	return
}

func GetUserByMobile(mobile string) (user model.User, err error) {
	if err = db.Conn().Get(&user, "select * from miaosha_user where mobile = ?", mobile); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Printf("dao.GetUserByMobile() failed, err: %v", err)
		}
	}
	return
}

func SaveUser(user model.User) (id int64, err error) {
	var result sql.Result
	if result, err = db.Conn().Exec("insert into miaosha_user(mobile) values(?)", user.Mobile); err != nil {
		log.Printf("dao.SaveUser() failed, err: %v", err)
		return
	}
	return result.LastInsertId()
}
