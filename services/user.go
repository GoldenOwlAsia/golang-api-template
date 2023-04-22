package services

import (
	"api/api/v1/requests"
	"api/api/v1/responses"
	"api/models"
	"api/pkgs/jwt_auth_token"
	"api/repository"
	"api/utils"
	"errors"
	"gorm.io/gorm"
	"os"
	"strconv"

	"strings"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		Repo: r,
	}
}

func (receiver UserService) Create(req requests.UserCreateRequest) (resp models.User, err error) {
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

func (receiver UserService) GetByUsername(username string) (resp models.User, err error) {
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
	var userIdString = strconv.FormatUint(uint64(userRes.Id), 10)
	if err = utils.VerifyPassword(userRes.Password, req.Password); err != nil {
		err = errors.New("invalid username or password")
		return
	}
	accessToken, _ := jwt_auth_token.GenerateAccessToken(userIdString, os.Getenv("SECRET_KEY"))
	refreshToken, _ := jwt_auth_token.GenerateRefreshToken(userIdString, os.Getenv("SECRET_KEY"))
	resp = responses.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return
}
