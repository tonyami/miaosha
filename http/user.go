package http

import (
	"github.com/gin-gonic/gin"
	"miaosha/conf"
	"miaosha/model"
	"net/http"
)

func userRegister(c *gin.Context) {
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	user := &model.User{
		Mobile:   mobile,
		Password: password,
	}
	if err := user.Check(); err != nil {
		c.JSON(http.StatusBadRequest, &Resp{
			Msg: err,
		})
		return
	}
	if err := userSrv.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, &Resp{
			Msg: err,
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func userLogin(c *gin.Context) {
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	user := &model.User{
		Mobile:   mobile,
		Password: password,
	}
	if err := user.Check(); err != nil {
		c.JSON(http.StatusBadRequest, &Resp{
			Msg: err,
		})
		return
	}
	token, err := userSrv.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Resp{
			Msg: err,
		})
		return
	}
	c.JSON(http.StatusOK, &Resp{
		Data: token,
	})
}

func userInfo(c *gin.Context) {
	user, _ := c.Get(conf.UserSession)
	c.JSON(http.StatusOK, &Resp{
		Data: user,
	})
}
