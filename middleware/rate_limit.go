package middleware

import (
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware(log *zap.Logger, requestsPerSec int, burstSize int) func(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(requestsPerSec), burstSize)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				log.Warn("rate limit exceeded", zap.String("path", r.URL.Path), zap.String("ip", r.RemoteAddr))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte(`{"error":"rate limit exceeded","code":429}`))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
