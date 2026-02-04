package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"am.com/gowebapp/cache"
	"am.com/gowebapp/logger"
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
	logger.Console("listing all users", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	// Try to get from cache
	if cacheService != nil {
		if cached, ok := cacheService.Get("users:list"); ok {
			logger.Console("users list retrieved from cache")
			writeSuccessWithCacheStatus(w, http.StatusOK, cached.(string), true)
			return
		}
	}

	// Not in cache, generate response
	message := "list all users"

	// Store in cache if cache service is available
	if cacheService != nil {
		cacheService.Set("users:list", message)
		logger.Console("users list stored in cache")
	}

	writeSuccessWithCacheStatus(w, http.StatusOK, message, false)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	logger.Console("creating new user", zap.String("method", r.Method), zap.String("path", r.URL.Path))
	writeSuccess(w, http.StatusCreated, "user created")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	logger.Console("fetching user", zap.String("userID", id), zap.String("method", r.Method), zap.String("path", r.URL.Path))

	cacheKey := fmt.Sprintf("user:%s", id)

	// Try to get from cache
	if cacheService != nil {
		if cached, ok := cacheService.Get(cacheKey); ok {
			logger.Console("user retrieved from cache", zap.String("userID", id))
			writeSuccessWithIDAndCacheStatus(w, http.StatusOK, cached.(string), id, true)
			return
		}
	}

	// Not in cache, generate response
	message := "get user"

	// Store in cache if cache service is available
	if cacheService != nil {
		cacheService.Set(cacheKey, message)
		logger.Console("user stored in cache", zap.String("userID", id))
	}

	writeSuccessWithIDAndCacheStatus(w, http.StatusOK, message, id, false)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	logger.Console("updating user", zap.String("userID", id), zap.String("method", r.Method), zap.String("path", r.URL.Path))
	writeSuccessWithID(w, http.StatusOK, "user updated", id)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	logger.ConsoleError("delete user operation failed", zap.String("userID", id), zap.String("reason", "cannot delete user"), zap.String("path", r.URL.Path))
	writeError(w, http.StatusBadRequest, "cannot delete user")
}
