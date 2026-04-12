package routes

import (
	"genderize-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	app.GET("/health", handlers.HealthCheck)

	api := app.Group("/api")
	api.GET("/classify", handlers.ClassifyHandler)

	return app
}