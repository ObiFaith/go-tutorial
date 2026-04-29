package mappers

import (
	"profile-api/internal/dtos"
	"profile-api/internal/models"
	"time"
)

func ToProfileResponse(p models.Profile) dtos.ProfileResponse {
	return dtos.ProfileResponse{
		Id:                p.ID,
		Name:              p.Name,
		Gender:            p.Gender,
		GenderProbability: p.GenderProbability,
		Age:               p.Age,
		AgeGroup:          p.AgeGroup,
		CountryId:         p.CountryId,
		CountryName:       p.CountryName,
		CountryProbability:p.CountryProbability,
		CreatedAt:         p.CreatedAt.UTC().Format(time.RFC3339),
	}
}

func ToProfileResponseList(profiles []models.Profile) []dtos.ProfileResponse {
	responses := make([]dtos.ProfileResponse, 0, len(profiles))

	for _, p := range profiles {
		responses = append(responses, ToProfileResponse(p))
	}

	return responses
}

