package handlers

import (
	"profile-api/internal/models"
	"profile-api/internal/services"
	"profile-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) CreateProfile(ctx *gin.Context) {
	var req models.CreateProfileRequest

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