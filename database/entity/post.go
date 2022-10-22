package entity

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID         uint64    `json:"id" gorm:"primaryKey;column:id"`
	UserID     uint64    `gorm:"type:uint;column:user_id;<-:create;not null"`
	User       User      `json:"user"`
	Content    string    `json:"content" binding:"required" gorm:"unique;type:string;size:777;column:content;<-:create;not null"`
	Repost     bool      `json:"repost" gorm:"type:boolean;column:repost;default:false"`
	Quote      []*Post   `json:"quote" gorm:"many2many:quote;association_jointable_foreignkey:quote_id"`
	Created_at time.Time `gorm:"autoCreateTime;not null"`
}
