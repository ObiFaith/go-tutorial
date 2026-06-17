package api

import (
	"genderize-api/internal/routes"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	once   sync.Once
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		router = routes.SetupRouter()
	})
	router.ServeHTTP(w, r)
}
