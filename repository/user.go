package repository

import (
	"posterr/database/entity"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserData(userUuid string) (entity.User, error)
}

type userRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepositoryInterface {
	return &userRepositoryStruct{DbConn: dbConn}
}

func (p *userRepositoryStruct) GetUserData(userUuid string) (entity.User, error) {
	userSearch := entity.User{Uuid: userUuid}

	p.DbConn.First(&userSearch)

	return userSearch, nil
}
