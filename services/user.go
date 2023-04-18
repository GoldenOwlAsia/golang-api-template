package services

import (
	"api/configs"
	"api/handler/api/v1/requests"
	"api/handler/api/v1/responses"
	"api/models/gorms"
	"api/pkgs/jwt"
	"api/repository"
	"api/utils"
	"errors"
	"gorm.io/gorm"

	// "fmt"
	"strings"
	"time"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		Repo: r,
	}
}

func (receiver UserService) Create(req requests.UserCreateRequest) (resp gorms.User, err error) {
	// valid data
	if req.Password != req.ConfirmPassword {
		err = errors.New("confirm password does not match password")
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		err = errors.New("have an error when create user")
		return
	}

	user, err := receiver.Repo.GetByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.New("have an error when get user")
		return
	}

	if len(user.Username) > 0 {
		err = errors.New("username already exists")
		return
	}

	req.Password = hashPassword
	req.Email = strings.ToLower(req.Email)

	resp, err = receiver.Repo.Create(req)
	if err != nil {
		err = errors.New("have an error when create user")
		return
	}
	return
}

func (receiver UserService) GetByUsername(username string) (resp gorms.User, err error) {
	resp, err = receiver.Repo.GetByUsername(username)
	if err != nil {
		err = errors.New("not found username")
		return
	}

	return
}

func (receiver UserService) Login(req requests.UserLoginRequest) (resp responses.UserLoginResponse, err error) {
	userRes, err := receiver.Repo.GetByUsername(req.Username)
	if err != nil || len(userRes.Username) <= 0 {
		err = errors.New("invalid username or password")
		return
	}

	if err = utils.VerifyPassword(userRes.Password, req.Password); err != nil {
		err = errors.New("invalid username or password")
		return
	}

	expiredTime := time.Now().Add(configs.ConfApp.TokenExpiresIn)
	token, err := jwt.GenerateToken(userRes, expiredTime, configs.ConfApp.TokenSecret)
	if err != nil {
		return
	}

	resp = responses.UserLoginResponse{
		Token:   token,
		Expires: expiredTime.Format("2006-01-02 15:04:05"),
	}
	return
}
