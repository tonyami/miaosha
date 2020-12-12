package user

import (
	"miaosha/conf"
	"miaosha/dao/user"
)

type Service struct {
	c *conf.Config
	d *user.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		c: c,
		d: user.New(c),
	}
}
