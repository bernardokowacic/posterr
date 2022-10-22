package repository

import (
	"posterr/database/entity"

	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	Get() error
	Insert(post entity.Post) error
}

type postRepositoryStruct struct {
	DbConn *gorm.DB
}

func NewPostRepository(dbConn *gorm.DB) PostRepositoryInterface {
	return &postRepositoryStruct{DbConn: dbConn}
}

func (p *postRepositoryStruct) Get() error {
	return nil
}

func (p *postRepositoryStruct) Insert(post entity.Post) error {

}
