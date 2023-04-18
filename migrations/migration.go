package migrations

import (
	"api/models"
	"api/pkgs/gorm"
	"log"
)

func Migrate() {
	log.Println("migrating data...")
	db := gorm.CreateInstanceDb()
	db.AutoMigrate(
		&models.User{},
	)
}
