package seeder

import (
	"log"
	"profile-api/internal/utils"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedProfiles(database *gorm.DB) error {
	start := time.Now()
	log.Println("Starting profile seeding...")

	profiles, err := utils.LoadProfiles("internal/database/seeder/profiles.json");

	if err != nil {
		log.Println("Failed to load profiles:", err)
		return err
	}

	log.Printf("Loaded %d profiles\n", len(profiles))

	// batch insert with conflict handling
	err = database.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}}, // unique field
		DoNothing: true,
	}).CreateInBatches(profiles, 100).Error

	if err != nil {
		return err
	}

	log.Printf("Seeding completed in %s\n", time.Since(start))

	return nil
}