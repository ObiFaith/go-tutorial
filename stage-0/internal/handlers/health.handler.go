package handlers

import (
	"genderize-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /health [get]
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.HealthResponse{Status: "ok"})
}
