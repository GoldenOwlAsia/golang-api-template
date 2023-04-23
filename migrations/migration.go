package migrations

import (
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) (err error) {
	return db.AutoMigrate(
		&models.User{},
		&models.Article{},
	)
}
