package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterUserV2Routes registers all user-related routes for v2 API
func RegisterUserV2Routes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", listUsersV2)
		r.Get("/{id}", getUserV2)
	})
}

// listUsersV2 returns a list of users with enhanced v2 response format
func listUsersV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"data":[],"pagination":{"total":0,"page":1},"message":"list all users - v2"}`))
}

// getUserV2 retrieves a user by ID with enhanced v2 response format
func getUserV2(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"data":{"id":"` + id + `","name":""},"message":"get user - v2"}`))
}
