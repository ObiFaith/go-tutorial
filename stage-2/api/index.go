package api

import (
	"net/http"
	"profile-api/config"
	"profile-api/internal/clients"
	"profile-api/internal/routes"
	"profile-api/internal/services"
)

var router http.Handler

func init() {
	cfg := config.LoadConfig()

	client := &clients.Client{
		GenderizeUrl:   cfg.GenderizeApi,
		AgifyUrl:       cfg.AgifyApi,
		NationalizeUrl: cfg.NationalizeApi,
	}

	profileService := services.NewProfileService(client)
	router = routes.SetupRouter(profileService)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
