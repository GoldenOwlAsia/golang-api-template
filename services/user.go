package services

import (
	"errors"
	"github.com/GoldenOwlAsia/golang-api-template/api/v1/requests"
	"github.com/GoldenOwlAsia/golang-api-template/api/v1/responses"
	"github.com/GoldenOwlAsia/golang-api-template/pkgs/jwt_auth_token"
	"github.com/GoldenOwlAsia/golang-api-template/repository"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"github.com/spf13/cast"
	"os"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{
		Repo: r,
	}
}

func (s UserService) Login(req requests.UserLoginRequest) (resp responses.UserLoginResponse, err error) {
	userRes, err := s.Repo.GetByUsername(req.Username)
	if err != nil || len(userRes.Username) <= 0 {
		err = errors.New("invalid username or password")
		return
	}
	var userIdString = cast.ToString(userRes.ID)
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
