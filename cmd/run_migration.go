package main

import (
	"api/configs"
	"api/migrations"
	"api/models"
	"api/pkgs/gorm"
	"api/utils"
	"flag"
	"log"
)

func main() {
	var seed bool
	flag.BoolVar(&seed, "seed", false, "seed database")
	flag.Parse()

	if seed {
		migrations.Migrate()
		RunSeed()
	} else {
		migrations.Migrate()
	}
}

func RunSeed() {
	if configs.ConfApp.AppMode == "release" {
		return
	}
	seedUser()
}

func seedUser() {
	db := gorm.CreateInstanceDb()
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
			log.Fatalf("failed to create user: %v", result.Error)
		}
		log.Printf("created user with ID: %d\n", user.Id)
	}
}
