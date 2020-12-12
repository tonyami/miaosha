package model

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Status     int8   `json:"-"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
