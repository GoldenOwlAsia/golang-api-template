package migrations

import (
	gormModels "api/models/gorms"
	"api/pkgs/gorm"
	"log"
)

func Migrate() {
	log.Println("migrating data...")
	db := gorm.CreateInstanceDb()
	db.AutoMigrate(
		&gormModels.User{},
	)
}
