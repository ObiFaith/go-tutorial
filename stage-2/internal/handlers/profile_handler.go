package handlers

import (
	"profile-api/internal/dtos"
	"profile-api/internal/services"
	"profile-api/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) CreateProfile(ctx *gin.Context) {
	var req dtos.CreateProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, "Missing or empty name parameter")
		return
	}

	if !utils.IsValidName(req.Name) {
		utils.UnprocessableEntity(ctx, "Invalid type")
		return
	}

	profile, err := h.service.CreateProfile(ctx.Request.Context(), req.Name)
	if err != nil {
		utils.BadGateway(ctx, err.Error())
		return
	}

	utils.CreatedResponse(ctx, profile)
}

func (h *ProfileHandler) GetProfiles(ctx *gin.Context) {
	gender := strings.ToLower(ctx.Query("gender"))
	minAge := strings.ToLower(ctx.Query("min_age"))
	maxAge := strings.ToLower(ctx.Query("max_age"))
	ageGroup := strings.ToLower(ctx.Query("age_group"))
	countryId := strings.ToUpper(ctx.Query("country_id"))
	minGenderProbability := strings.ToLower(ctx.Query("min_gender_probability"))
	minCountryProbability := strings.ToLower(ctx.Query("min_country_probability"))

	filters := dtos.ProfileFilter{
		MinAge:    						 minAge,
		MaxAge:								 maxAge,
		Gender:    						 gender,
		AgeGroup:  						 ageGroup,
		CountryId: 						 countryId,
		MinGenderProbability:  minGenderProbability,
		MinCountryProbability: minCountryProbability,
	}

	profiles, err := h.service.GetProfiles(ctx.Request.Context(), filters)
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.OKResponse(ctx, profiles)
}

func (h *ProfileHandler) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		utils.BadRequest(ctx, "Missing id parameter")
		return
	}

	profile, err := h.service.GetProfile(ctx.Request.Context(), id)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	utils.OKResponse(ctx, profile)
}

func (h *ProfileHandler) DeleteProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		utils.BadRequest(ctx, "Missing id parameter")
		return
	}

	err := h.service.DeleteProfile(ctx.Request.Context(), id)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	utils.NoContent(ctx)
}
