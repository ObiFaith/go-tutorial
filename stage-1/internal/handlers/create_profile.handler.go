package handlers

import (
	"net/http"
	"strings"
	"time"

	"genderize-api/internal/models"
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
		utils.BadRequest(ctx, "invalid JSON payload")
		return
	}

	profile.Name = strings.ToLower(strings.TrimSpace(profile.Name))

	if profile.Name == "" {
		utils.BadRequest(ctx, "name is required")
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

	response := models.ProfileResponse{
		Name:               genderRes.Name,
		Gender:             genderRes.Gender,
		GenderProbability:  genderRes.Probability,
		SampleSize:         genderRes.Count,
		Age:                ageRes.Age,
		AgeGroup:           utils.AgeGroup(ageRes.Age),
		CountryID:          countryRes.CountryID,
		CountryProbability: countryRes.Probability,
		CreatedAt:          time.Now().UTC().Format(time.RFC3339),
	}

	utils.CreatedResponse(ctx, response)
}