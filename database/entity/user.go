package entity

import (
	"time"
)

type User struct {
	ID         uint64    `json:"id" gorm:"primaryKey;column:id"`
	Username   string    `json:"username" binding:"required" gorm:"unique;type:string;size:14;column:username;not null"`
	Uuid       string    `json:"uuid" gorm:"unique;type:string;size:36;column:uuid;not null"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	TotalPosts uint64    `json:"total_posts"`
}
