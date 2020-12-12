package user

import (
	"miaosha/conf"
)

var s *Service

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	s = New(conf.Conf)
}
