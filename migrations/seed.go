package migrations

import (
	"github.com/GoldenOwlAsia/golang-api-template/configs"
	"github.com/GoldenOwlAsia/golang-api-template/models"
	"github.com/GoldenOwlAsia/golang-api-template/utils"
	"gorm.io/gorm"
	"log"
)

func Seed(DB *gorm.DB) {
	if configs.ConfApp.AppMode == "release" {
		return
	}
	log.Println("Seeding Data...")
	seedUser(DB)
}

func seedUser(DB *gorm.DB) {
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
		result := DB.Where("username = ?", user.Username).FirstOrCreate(&user)
		if result.Error != nil {
			log.Printf("failed to create user: %v", result.Error)
		}
		log.Printf("created user with ID: %d\n", user.ID)
	}
}
