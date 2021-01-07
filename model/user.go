package model

import "time"

type User struct {
	Id         int64     `json:"id"`
	Mobile     string    `json:"mobile"`
	CreateTime time.Time `json:"create_time"`
}
