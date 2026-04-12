package handlers

import (
	"net/http"
	"strings"

	"genderize-api/internal/services"

	"github.com/gin-gonic/gin"
)

func ClassifyHandler(ctx *gin.Context) {
	name := strings.TrimSpace(ctx.Query("name"))
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "missing or empty name parameter",
		})
		return
	}

	data, err := services.FetchGenderData(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Upstream or server failure",
		})
		return
	}

	if (data.Gender == "" || data.Count == 0){
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "No prediction available for the provided name",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"name":         data.Name,
			"gender":       data.Gender,
			"probability":  data.Probability,
			"sample_size":  data.Count,
			"is_confident": data.Probability >= 0.7 && data.Count >= 100,
		},
	})
}