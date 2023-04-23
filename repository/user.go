package repository

import (
	"github.com/GoldenOwlAsia/golang-api-template/api/v1/requests"
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"github.com/getsentry/sentry-go"
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

func (receiver UserRepository) Create(req requests.UserCreateRequest) (resp models.User, err error) {
	userGorm := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     configs.UserRoleAdmin,
		Status:   configs.UserStatusActive,
	}
	tx := receiver.DB.Create(&userGorm)

	err = tx.Error
	if tx.Error != nil {
		sentry.CaptureException(tx.Error)
		return
	}

	resp = userGorm
	return
}

func (receiver UserRepository) GetByUsername(username string) (resp models.User, err error) {
	err = receiver.DB.Where(&models.User{Username: username}).First(&resp).Error
	return
}
