package model

import "time"

type User struct {
	Id           int64     `json:"id"`
	Mobile       string    `json:"mobile"`
	RegisterTime time.Time `json:"register_time"`
}
