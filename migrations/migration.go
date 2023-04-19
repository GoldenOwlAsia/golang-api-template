package migrations

import (
	"api/api"
	"api/models"
	"log"
)

func Migrate() (err error) {
	log.Println("migrating data...")
	return api.Api.DB.AutoMigrate(
		&models.User{},
	)
}
