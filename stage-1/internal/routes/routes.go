package routes

import (
	"profile-api/internal/handlers"
	"profile-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(profileService *services.ProfileService) *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false, // TODO: change to true with strict origin
	}))

	app.GET("/health", handlers.HealthCheck)
	api := app.Group("/api")

	profileHandler := handlers.NewProfileHandler(profileService)
	api.GET("/profiles", profileHandler.GetProfiles)
	api.GET("/profiles/:id", profileHandler.GetProfile)
	api.POST("/profiles", profileHandler.CreateProfile)
	api.DELETE("/profiles/:id", profileHandler.DeleteProfile)

	return app
}