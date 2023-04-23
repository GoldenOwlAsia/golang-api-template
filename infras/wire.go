//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"github.com/GoldenOwlAsia/golang-api-template/handlers"
	"github.com/GoldenOwlAsia/golang-api-template/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) handlers.UserHandler {
	wire.Build(
		services.NewUserService,
		handlers.NewUserHandler,
	)

	return handlers.UserHandler{}
}

func InitArticleAPI(db *gorm.DB) handlers.ArticleHandler {
	wire.Build(
		services.NewArticleService,
		handlers.NewArticleHandler,
	)

	return handlers.ArticleHandler{}
}
