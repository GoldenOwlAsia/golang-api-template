package migrations

import (
	"log"

	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	if configs.ConfApp.AppMode == "release" {
		return
	}
	seedUser(db)
}

func seedUser(db *gorm.DB) {
	// Create some sample users
	hashPassword, err := utils.HashPassword("1234")
	if err != nil {
		log.Fatal(err)
	}
	users := []models.User{
		{Username: "admin", Email: "admin@example.com", Password: hashPassword, Role: "Admin"},
	}

	// Insert the users into the database
	for _, user := range users {
		result := db.Where("username = ?", user.Username).FirstOrCreate(&user)
		if result.Error != nil {
			log.Printf("failed to create user: %v", result.Error)
		}
		log.Printf("created user with ID: %d\n", user.ID)
	}
}
