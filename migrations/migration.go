package migrations

import (
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(DB *gorm.DB) (err error) {
	log.Println("migrating data...")
	return DB.AutoMigrate(
		&models.User{},
		&models.Article{},
	)
}
