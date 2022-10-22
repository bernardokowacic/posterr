package database

import (
	"posterr/database/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Migrator().CreateTable(&entity.User{})
	db.Migrator().CreateTable(&entity.Post{})
}
