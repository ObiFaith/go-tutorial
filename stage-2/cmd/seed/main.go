package main

import (
	"log"
	"profile-api/config"
	"profile-api/internal/database"
	"profile-api/internal/database/seeder"
)

func main(){
	cfg := config.LoadConfig()

	database.Connect(cfg.DatabaseUrl)
	database.Migrate()

	if err := seeder.SeedProfiles(database.DB); err != nil {
		log.Fatal("seeding failed:", err)
	}

	log.Println("seeding completed successfully")
}