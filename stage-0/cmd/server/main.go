package main

import (
	"genderize-api/config"
	"genderize-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Genderize API
// @version 1.0
// @description This is a sample server for Genderize API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
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
