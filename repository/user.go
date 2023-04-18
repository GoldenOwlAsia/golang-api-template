package repository

import (
	"api/configs"
	"api/handler/api/v1/requests"
	"api/models/gorms"
	"time"

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

func (receiver UserRepository) Create(req requests.UserCreateRequest) (resp gorms.User, err error) {
	defaultRole := configs.UserRoleAdmin
	defaultStatus := configs.UserStatusActive
	defaultApprovedStatus := configs.UserApprovedStatus
	userGorm := gorms.User{
		Username:       req.Username,
		Password:       req.Password,
		Email:          req.Email,
		Role:           defaultRole,
		Status:         defaultStatus,
		ApprovedStatus: defaultApprovedStatus,
		CreatedBy:      req.Username,
		UpdatedBy:      req.Username,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
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

func (receiver UserRepository) GetByUsername(username string) (resp gorms.User, err error) {
	tx := receiver.DB.Where(&gorms.User{Username: username}).First(&resp)

	err = tx.Error
	if err != nil {
		sentry.CaptureException(tx.Error)
	}
	return
}
