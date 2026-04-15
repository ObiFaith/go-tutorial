package api

import (
	"genderize-api/internal/routes"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := routes.SetupRouter()
	router.ServeHTTP(w, r)
}