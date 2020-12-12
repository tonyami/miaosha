package user

import "miaosha/conf"

var d *Dao

func init() {
	if err := conf.Init("test"); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
}
