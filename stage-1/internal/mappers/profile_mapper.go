package mappers

import (
	"profile-api/internal/dtos"
	"profile-api/internal/models"
	"time"
)

func ToProfileResponse(p models.Profile) *dtos.ProfileResponse {
	return &dtos.ProfileResponse{
		Id:                p.ID,
		Name:              p.Name,
		Gender:            p.Gender,
		GenderProbability: p.GenderProbability,
		SampleSize:        p.SampleSize,
		Age:               p.Age,
		AgeGroup:          p.AgeGroup,
		CountryID:         p.CountryID,
		CountryProbability:p.CountryProbability,
		CreatedAt:         p.CreatedAt.Format(time.RFC3339),
	}
}

