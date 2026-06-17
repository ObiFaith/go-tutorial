package handlers

import (
	"genderize-api/internal/models"
	"genderize-api/internal/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ClassifyHandler godoc
// @Summary Classify name
// @Description Get gender classification for a name
// @Tags classify
// @Accept json
// @Produce json
// @Param name query string true "Name to classify"
// @Success 200 {object} models.ClassifyResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /classify [get]
func ClassifyHandler(ctx *gin.Context) {
	name := strings.TrimSpace(ctx.Query("name"))
	if name == "" {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:  "error",
			Message: "missing or empty name parameter",
		})
		return
	}

	data, err := services.FetchGenderData(ctx.Request.Context(), name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:  "error",
			Message: "Upstream or server failure",
		})
		return
	}

	if (data.Gender == "" || data.Count == 0) {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			Status:  "error",
			Message: "No prediction available for the provided name",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ClassifyResponse{
		Status: "success",
		Data: models.ClassifyResponseData{
			Name:        data.Name,
			Gender:      data.Gender,
			Probability: data.Probability,
			SampleSize:  data.Count,
			IsConfident: data.Probability >= 0.7 && data.Count >= 100,
			ProcessedAt: time.Now().UTC().Format(time.RFC3339),
		},
	})
}
