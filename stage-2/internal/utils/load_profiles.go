package utils

import (
	"encoding/json"
	"os"
	"profile-api/internal/models"
)

type ProfileWrapper struct {
	Profiles []models.Profile `json:"profiles"`
}

func LoadProfiles(path string) ([]models.Profile, error){
	var data ProfileWrapper

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	return data.Profiles, err
}