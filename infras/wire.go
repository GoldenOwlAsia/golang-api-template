//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	controllers2 "github.com/GoldenOwlAsia/golang-api-template/handlers"
	"github.com/GoldenOwlAsia/golang-api-template/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) controllers2.UserHandler {
	wire.Build(
		services.NewUserService,
		controllers2.NewUserHandler,
	)

	return controllers2.UserHandler{}
}

func InitArticleAPI(db *gorm.DB) controllers2.ArticleHandler {
	wire.Build(
		services.NewArticleService,
		controllers2.NewArticleHandler,
	)

	return controllers2.ArticleHandler{}
}
