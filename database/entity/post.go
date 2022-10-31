package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID         uint64    `json:"id" gorm:"primaryKey;column:id"`
	UserID     uint64    `json:"-" gorm:"type:uint;column:user_id;<-:create;not null"`
	User       User      `json:"user"`
	Content    string    `json:"content" binding:"required" gorm:"type:string;size:777;column:content;<-:create;not null"`
	RepostID   *uint64   `json:"repost" gorm:"type:uint;column:repost_id"`
	QuoteID    *uint64   `json:"quote" gorm:"type:uint;column:quote_id"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
}
