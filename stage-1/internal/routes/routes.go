package routes

import (
	"genderize-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
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
	api.POST("/profiles", handlers.CreateProfileHandler)

	return app
}