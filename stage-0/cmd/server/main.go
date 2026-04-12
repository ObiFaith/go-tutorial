package main

import (
	"genderize-api/config"
	"genderize-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.GinMode == "release" {
    gin.SetMode(gin.ReleaseMode)
	} else {
			gin.SetMode(gin.DebugMode)
	}

	app := routes.SetupRouter()

	if err := app.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}