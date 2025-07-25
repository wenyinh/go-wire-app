package entity

import "time"

type UserEntity struct {
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Gender     string     `json:"gender"`
	Age        int        `json:"age"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}
