package api

import (
	"net/http"
	"profile-api/internal/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := routes.SetupRouter()
	router.ServeHTTP(w, r)
}