package entity

import (
	"time"
)

type Post struct {
	ID         uint64    `json:"id" gorm:"primaryKey;column:id"`
	UserID     uint64    `json:"user_id" gorm:"type:uint;column:user_id;<-:create;not null"`
	User       User      `json:"-"`
	Content    string    `json:"content" binding:"required" gorm:"type:string;size:777;column:content;<-:create;not null"`
	RepostID   *uint64   `json:"repost_id" gorm:"type:uint;column:repost_id"`
	QuoteID    *uint64   `json:"quote_id" gorm:"type:uint;column:quote_id"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	Repost     *Post     `json:"repost"`
	Quote      *Post     `json:"quote"`
}
