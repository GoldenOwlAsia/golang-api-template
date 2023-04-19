//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"api/api/v1"
	"api/repository"
	"api/services"
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
