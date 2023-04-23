//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"github.com/GoldenOwlAsia/golang-api-template/api/v1"
	"github.com/GoldenOwlAsia/golang-api-template/repository"
	"github.com/GoldenOwlAsia/golang-api-template/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) v1.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		services.NewUserService,
		v1.NewUserHandler,
	)

	return v1.UserHandler{}
}

func InitArticleAPI(db *gorm.DB) v1.ArticleHandler {
	wire.Build(
		v1.NewArticleHandler,
	)

	return v1.ArticleHandler{}
}
