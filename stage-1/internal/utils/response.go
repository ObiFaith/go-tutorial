package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"message": message,
	})
}

func UnprocessableEntity(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"status": "error",
		"message": message,
	})
}

func BadGateway(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadGateway, gin.H{
		"status": "error",
		"message": message,
	})
}

func InternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status": "error",
		"message": message,
	})
}

func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"status": "error",
		"message": message,
	})
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}


func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"status": "error",
		"message": message,
	})
}

func OKResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

func CreatedResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   data,
	})
}