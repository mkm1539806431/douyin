package model

import "time"

type TUser struct {
	ID            int64 `gorm:"primaryKey autoIncrement"`
	Username      string
	Nickname      string
	Password      string
	Token         string
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
