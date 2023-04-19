package migrations

import (
	"api/api"
	"api/configs"
	"api/models"
	"api/utils"
	"log"
)

func Seed() {
	if configs.ConfApp.AppMode == "release" {
		return
	}
	log.Println("Seeding Data...")
	seedUser()
}

func seedUser() {
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
		result := api.Api.DB.Where("username = ?", user.Username).FirstOrCreate(&user)
		if result.Error != nil {
			log.Fatalf("failed to create user: %v", result.Error)
		}
		log.Printf("created user with ID: %d\n", user.Id)
	}
}
