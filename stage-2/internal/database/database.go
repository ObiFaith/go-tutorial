package database

import (
	"log"
	"profile-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string){
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	DB = database
	log.Println("Database connected")
}

func Migrate() {
	log.Println("Running migrations...")
	err := DB.AutoMigrate(&models.Profile{})
	if err != nil {
		log.Fatal("migration failed:", err)
	}
}
