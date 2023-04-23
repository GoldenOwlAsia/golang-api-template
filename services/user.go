package services

import (
	"errors"
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/handlers/requests"
	"github.com/GoldenOwlAsia/golang-api-template/handlers/responses"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	jwt_token "github.com/GoldenOwlAsia/golang-api-template/pkgs/jwt-token"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{
		DB: db,
	}
}

func (s UserService) Login(req requests.UserLoginRequest) (resp responses.UserLoginResponse, err error) {
	var user models.User
	err = s.DB.Where(&models.User{Username: req.Username}).First(&user).Error

	if err != nil || len(user.Username) <= 0 {
		err = errors.New("invalid username or password")
		return
	}
	var userIDString = cast.ToString(user.ID)
	if err = utils.VerifyPassword(user.Password, req.Password); err != nil {
		err = errors.New("invalid username or password")
		return
	}
	accessToken, _ := jwt_token.GenerateAccessToken(userIDString, configs.ConfApp.SecretKey)
	refreshToken, _ := jwt_token.GenerateRefreshToken(userIDString, configs.ConfApp.SecretKey)
	resp = responses.UserLoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return
}
