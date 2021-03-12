package model

import "time"

type User struct {
	Id         int64     `db:"id" json:"id"`
	Mobile     string    `db:"mobile" json:"mobile"`
	CreateTime time.Time `db:"create_time" json:"-"`
	UpdateTime time.Time `db:"update_time" json:"-"`
}
