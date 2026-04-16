package services

import (
	"context"
	"genderize-api/config"
	"genderize-api/internal/clients"
	"genderize-api/internal/models"

	"golang.org/x/sync/errgroup"
)

type ProfileService struct {
	cfg config.Config
}

func (s *ProfileService) CreateProfile(ctx context.Context, name string) (*models.ProfileResponse, error) {
	var (
		gender  clients.GenderResponse
		age     clients.AgeResponse
		country clients.Country
	)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		res, err := clients.FetchGender(ctx, s.cfg.GenderizeApi, name)
		if err != nil {
			return err
		}
		gender = res
		return nil
	})
	group.Go(func() error {
		res, err := clients.FetchAge(ctx, s.cfg.AgifyApi, name)
		if err != nil {
			return err
		}
		age = res
		return nil
	})

}