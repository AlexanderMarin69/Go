package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterUserRoutes registers all user-related routes for v1 API
func RegisterUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", listUsers)
		r.Post("/", createUser)
		r.Get("/{id}", getUser)
		r.Put("/{id}", updateUser)
		r.Delete("/{id}", deleteUser)
	})
}

// listUsers returns a list of all users
func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"list all users"}`))
}

// createUser creates a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"user created"}`))
}

// getUser retrieves a user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"get user","id":"` + id + `"}`))
}

// updateUser updates a user by ID
func updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"user updated","id":"` + id + `"}`))
}

// deleteUser deletes a user by ID - returns bad request for testing
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"cannot delete user","code":"DELETE_FORBIDDEN"}`))
}
