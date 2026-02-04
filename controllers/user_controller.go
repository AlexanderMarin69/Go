package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"am.com/gowebapp/cache"
)

var cacheService *cache.Service

// RegisterUserRoutes registers all user-related routes for v1 API
func RegisterUserRoutes(r chi.Router, cs *cache.Service) {
	cacheService = cs

	r.Route("/users", func(r chi.Router) {
		r.Get("/", listUsers)
		r.Post("/", createUser)
		r.Get("/{id}", getUser)
		r.Put("/{id}", updateUser)
		r.Delete("/{id}", deleteUser)
	})
}

// Handlers
func listUsers(w http.ResponseWriter, r *http.Request) {
	// Try to get from cache
	if cacheService != nil {
		if cached, ok := cacheService.Get("users:list"); ok {
			writeSuccessWithCacheStatus(w, http.StatusOK, cached.(string), true)
			return
		}
	}

	// Not in cache, generate response
	message := "list all users"

	// Store in cache if cache service is available
	if cacheService != nil {
		cacheService.Set("users:list", message)
	}

	writeSuccessWithCacheStatus(w, http.StatusOK, message, false)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	writeSuccess(w, http.StatusCreated, "user created")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	cacheKey := fmt.Sprintf("user:%s", id)

	// Try to get from cache
	if cacheService != nil {
		if cached, ok := cacheService.Get(cacheKey); ok {
			writeSuccessWithIDAndCacheStatus(w, http.StatusOK, cached.(string), id, true)
			return
		}
	}

	// Not in cache, generate response
	message := "get user"

	// Store in cache if cache service is available
	if cacheService != nil {
		cacheService.Set(cacheKey, message)
	}

	writeSuccessWithIDAndCacheStatus(w, http.StatusOK, message, id, false)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writeSuccessWithID(w, http.StatusOK, "user updated", id)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusBadRequest, "cannot delete user")
}
