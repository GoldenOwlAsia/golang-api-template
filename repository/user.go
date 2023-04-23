package repository

import (
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (r UserRepository) GetByUsername(username string) (resp models.User, err error) {
	err = r.DB.Where(&models.User{Username: username}).First(&resp).Error
	return
}
