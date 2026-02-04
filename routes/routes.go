package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"am.com/gowebapp/config"
	"am.com/gowebapp/controllers"
	"am.com/gowebapp/middleware"
)

// NewRouter creates and configures the main router with all routes
func NewRouter(log *zap.Logger, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.LoggingMiddleware(log))

	// Rate limiting middleware (if enabled)
	if cfg.RateLimit.Enabled {
		r.Use(middleware.RateLimitMiddleware(log, cfg.RateLimit.RequestsPerSec, cfg.RateLimit.BurstSize))
	}

	// Response caching middleware (if enabled)
	if cfg.Cache.Enabled {
		cache := middleware.NewResponseCache(cfg.Cache.MaxSize, cfg.Cache.TTL, log)
		r.Use(middleware.CachingMiddleware(cache, log))
	}

	// Health check endpoint
	r.Get("/health", healthCheck)

	// API v1 routes
	v1 := chi.NewRouter()
	controllers.RegisterUserRoutes(v1)
	r.Mount("/api/v1", v1)

	return r
}

// healthCheck is a simple health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
