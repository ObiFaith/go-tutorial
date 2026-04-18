package services

import (
	"context"
	"fmt"
	"profile-api/internal/clients"
	"profile-api/internal/dtos"
	"profile-api/internal/mappers"
	"profile-api/internal/models"
	"profile-api/internal/utils"
	"time"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type ProfileService struct {
	client 		*clients.Client
	database  *gorm.DB
}

func NewProfileService(client *clients.Client, database  *gorm.DB) *ProfileService {
	return &ProfileService{
		client: 	client,
		database: database,
	}
}

func (s *ProfileService) CreateProfile(ctx context.Context, name string) (*dtos.ProfileResponse, error) {
	var existing models.Profile

	err := s.database.WithContext(ctx).Where("name = ?", name).First(&existing).Error

	if err == nil {
		return mappers.ToProfileResponse(existing), nil
	}

	var (
		gender  clients.GenderResponse
		age     clients.AgeResponse
		country clients.Country
	)

	group, apiCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		res, err := s.client.FetchGender(apiCtx, name)
		if err != nil {
			return fmt.Errorf("Genderize returned an invalid response")
		}
		gender = res
		return nil
	})

	group.Go(func() error {
		res, err := s.client.FetchAge(apiCtx, name)
		if err != nil {
			return fmt.Errorf("Agify returned an invalid response")
		}
		age = res
		return nil
	})

	group.Go(func() error {
		res, err := s.client.FetchCountry(apiCtx, name)
		if err != nil {
			return fmt.Errorf("Nationalize returned an invalid response")
		}
		country = res
		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}

	if gender.Gender == "" || gender.Count == 0 {
		return nil, fmt.Errorf("No prediction available for the provided name")
	}

	profile := models.Profile{
		Name:               gender.Name,
		Gender:             gender.Gender,
		GenderProbability:  gender.Probability,
		SampleSize:         gender.Count,
		Age:                age.Age,
		AgeGroup:           utils.AgeGroup(age.Age),
		CountryID:          country.CountryID,
		CountryProbability: country.Probability,
	}

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.database.WithContext(dbCtx).Create(&profile).Error; err != nil {
		return nil, fmt.Errorf("Failed to create profile: %w", err)
	}

	return mappers.ToProfileResponse(profile), nil
}

// func (s *ProfileService) GetProfile(ctx context.Context, name string) (*dtos.ProfileResponse, error){
	
// }