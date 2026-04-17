package main

import (
	"log"
	"profile-api/config"
	"profile-api/internal/clients"
	"profile-api/internal/routes"
	"profile-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.GinMode == "release" {
    gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	client := &clients.Client{
		GenderizeUrl:   cfg.GenderizeApi,
		AgifyUrl:       cfg.AgifyApi,
		NationalizeUrl: cfg.NationalizeApi,
	}

	profileService := services.NewProfileService(client)

	app := routes.SetupRouter(profileService)

	if err := app.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}