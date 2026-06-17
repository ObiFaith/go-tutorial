package routes

import (
	"genderize-api/config"
	"genderize-api/docs"
	"genderize-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	cfg := config.LoadConfig()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false, // TODO: change to true with strict origin
	}))

	// Swagger documentation
	docs.SwaggerInfo.BasePath = "/api"
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := app.Group("/api")
	api.GET("/classify", handlers.ClassifyHandler)
	api.GET("/health", handlers.HealthCheck)

	return app
}
