package main

import (
	"genderize-api/config"
	"genderize-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.MODE == "release" {
    gin.SetMode(gin.ReleaseMode)
	} else {
			gin.SetMode(gin.DebugMode)
	}

	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	routes.RegisterRoutes(app)

	if err := app.Run(":" + cfg.PORT); err != nil {
		log.Fatal(err)
	}
}