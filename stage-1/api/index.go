package api

import (
	"net/http"
	"profile-api/config"
	"profile-api/internal/clients"
	"profile-api/internal/database"
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

	database.Connect(cfg.DatabaseUrl)
	database.Migrate()

	profileService := services.NewProfileService(client, database.DB)
	router = routes.SetupRouter(profileService)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
