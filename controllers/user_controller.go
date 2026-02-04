package controllers

import (
	"net/http"
	"strconv"

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

// Helper functions
func writeResponse(w http.ResponseWriter, statusCode int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(body))
}

func writeSuccess(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, `{"message":"`+message+`"}`)
}

func writeSuccessWithID(w http.ResponseWriter, statusCode int, message string, id string) {
	writeResponse(w, statusCode, `{"message":"`+message+`","id":"`+id+`"}`)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, `{"error":"`+message+`","code":`+strconv.Itoa(statusCode)+`}`)
}

// Handlers
func listUsers(w http.ResponseWriter, r *http.Request) {
	writeSuccess(w, http.StatusOK, "list all users")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	writeSuccess(w, http.StatusCreated, "user created")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writeSuccessWithID(w, http.StatusOK, "get user", id)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writeSuccessWithID(w, http.StatusOK, "user updated", id)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusBadRequest, "cannot delete user")
}
