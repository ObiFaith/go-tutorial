package services

import (
	"context"
	"errors"
	"fmt"
	"profile-api/internal/clients"
	"profile-api/internal/dtos"
	"profile-api/internal/mappers"
	"profile-api/internal/models"
	"profile-api/internal/utils"
	"strings"
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

func (s *ProfileService) CreateProfile(ctx context.Context, reqName string) (*dtos.ProfileResponse, error) {
	var existing models.Profile
	name := strings.ToLower(reqName)

	err := s.database.WithContext(ctx).Where("name = ?", name).First(&existing).Error

	if err == nil {
		resp := mappers.ToProfileResponse(existing)
		return &resp, nil
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
		Name:               name,
		Gender:             gender.Gender,
		GenderProbability:  gender.Probability,
		Age:                age.Age,
		AgeGroup:           utils.AgeGroup(age.Age),
		CountryId:          country.CountryId, // TODO: Add CountryName
		CountryProbability: country.Probability,
	}

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.database.WithContext(dbCtx).Create(&profile).Error; err != nil {
		return nil, fmt.Errorf("Failed to create profile: %w", err)
	}

	resp := mappers.ToProfileResponse(profile)
	return &resp, nil
}

func (s *ProfileService) GetProfiles(ctx context.Context, f dtos.ProfileFilter) ([]dtos.ProfileResponse, error){
	query := s.database.WithContext(ctx).Model(&models.Profile{})

	if f.Gender != "" {
		query = query.Where("LOWER(gender) = ?", f.Gender)
	}

	if f.CountryId != "" {
		query = query.Where("UPPER(country_id) = ?", f.CountryId)
	}

	if f.AgeGroup != "" {
		query = query.Where("LOWER(age_group) = ?", f.AgeGroup)
	}

	if f.MinAge != "" {
		query = query.Where("age >= ?", f.MinAge)
	}

	if f.MaxAge != "" {
		query = query.Where("age <= ?", f.MaxAge)
	}

	if f.MinCountryProbability != "" {
		query = query.Where("country_probability >= ?", f.MinCountryProbability)
	}

	if f.MinGenderProbability != "" {
		query = query.Where("gender_probability >= ?", f.MinGenderProbability)
	}

	var profiles []models.Profile

	if err := query.Find(&profiles).Error; err != nil {
		return nil, err
	}

	return mappers.ToProfileResponseList(profiles), nil
}

func (s *ProfileService) GetProfile(ctx context.Context, id string) (*dtos.ProfileResponse, error) {
	var existing models.Profile

	err := s.database.WithContext(ctx).
		Where("id = ?", id).
		First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("profile not found")
		}
		return nil, err
	}

	resp := mappers.ToProfileResponse(existing)
	return &resp, nil
}

func (s *ProfileService) DeleteProfile(ctx context.Context, id string) error{
	var existing models.Profile

	err := s.database.WithContext(ctx).Where("id = ?", id).First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return fmt.Errorf("Profile not found")
		}
		return err
	}

	if err := s.database.WithContext(ctx).Delete(&existing).Error; err != nil {
		return err
	}

	return nil
}