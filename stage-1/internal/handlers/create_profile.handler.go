package handlers

import (
	"net/http"
	"strings"
	"time"

	"genderize-api/internal/services"
	"genderize-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	Name string `json:"name" binding:"required"`
}

func CreateProfileHandler(ctx *gin.Context) {
	var profile Profile

	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid JSON payload",
		})
		return
	}

	profile.Name = strings.ToLower(strings.TrimSpace(profile.Name))

	if profile.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "name is required",
		})
		return
	}

	genderRes, genderErr := services.FetchGenderResponse(profile.Name)

	if genderErr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": "Genderize API returned an invalid response",
		})
		return
	}

	if (genderRes.Gender == "" || genderRes.Count == 0){
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": "Unable to determine gender for the provided name",
		})
		return
	}

	ageRes, ageErr := services.FetchAgeResponse(profile.Name)

	if ageErr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": "Agify API returned an invalid response",
		})
		return
	}

	if ageRes.Age == 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": "Unable to estimate age for the provided name",
		})
		return
	}

	countryRes, countryErr := services.FetchCountryResponse(profile.Name)

	if countryErr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": "Nationalize API returned an invalid response",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"name":         genderRes.Name,
			"gender":       genderRes.Gender,
			"gender_probability":  genderRes.Probability,
			"sample_size":  genderRes.Count,
			"age":  ageRes.Age,
			"age_group": utils.GetAgeGroup(ageRes.Age),
			"country_id": countryRes.CountryID,
			"country_probability": countryRes.Probability,
			"created_at": time.Now().UTC().Format(time.RFC3339),
		},
	})
}