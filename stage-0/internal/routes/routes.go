package routes

import (
	"genderize-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Count int
	Name string
	Gender string
	Probability float64
}

func RegisterRoutes (router *gin.Engine){
	router.GET("/health", handlers.HealthCheck)

	api := router.Group("/api")
	api.GET("/classify", handlers.ClassifyHandler)
}