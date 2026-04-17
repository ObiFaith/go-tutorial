package services

import (
	"context"
	"fmt"
	"profile-api/internal/clients"
	"profile-api/internal/models"
	"profile-api/internal/utils"
	"time"

	"golang.org/x/sync/errgroup"
)

type ProfileService struct {
	client *clients.Client
}

func NewProfileService(client *clients.Client) *ProfileService {
	return &ProfileService{client: client}
}

func (s *ProfileService) CreateProfile(ctx context.Context, name string) (*models.ProfileResponse, error) {
	var (
		gender  clients.GenderResponse
		age     clients.AgeResponse
		country clients.Country
	)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		res, err := s.client.FetchGender(ctx, name)
		if err != nil {
		return fmt.Errorf("Genderize returned an invalid response")
		}
		gender = res
		return nil
	})
	group.Go(func() error {
		res, err := s.client.FetchAge(ctx, name)
		if err != nil {
			return fmt.Errorf("Agify returned an invalid response")
		}
		age = res
		return nil
	})
	group.Go(func() error {
		res, err := s.client.FetchCountry(ctx, name)
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

	response := &models.ProfileResponse{
		Name:               gender.Name,
		Gender:             gender.Gender,
		GenderProbability:  gender.Probability,
		SampleSize:         gender.Count,
		Age:                age.Age,
		AgeGroup:           utils.AgeGroup(age.Age),
		CountryID:          country.CountryID,
		CountryProbability: country.Probability,
		CreatedAt:          time.Now().UTC().Format(time.RFC3339),
	}

	return response, nil
}