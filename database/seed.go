package database

import (
	"fmt"
	"posterr/database/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var users []entity.User

	for i := 0; i < 4; i++ {
		myuuid := uuid.New()
		username := fmt.Sprintf("User %d", i)

		user := entity.User{
			Username: username,
			Uuid:     myuuid.String(),
		}
		users = append(users, user)
	}

	db.Create(&users)
}
