package middleware

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

// CacheEntry represents a cached response
type CacheEntry struct {
	Body       []byte
	Headers    http.Header
	StatusCode int
	ExpiresAt  time.Time
}

// ResponseCache manages response caching with TTL and size limits
type ResponseCache struct {
	mu       sync.RWMutex
	cache    map[string]*CacheEntry
	maxSize  int
	ttl      time.Duration
	log      *zap.Logger
}

// NewResponseCache creates a new response cache
func NewResponseCache(maxSize int, ttl time.Duration, log *zap.Logger) *ResponseCache {
	rc := &ResponseCache{
		cache:   make(map[string]*CacheEntry),
		maxSize: maxSize,
		ttl:     ttl,
		log:     log,
	}

	// Start cleanup goroutine for expired entries
	go rc.cleanupExpiredEntries()

	return rc
}

// generateCacheKey creates a cache key from request method and URL
func (rc *ResponseCache) generateCacheKey(r *http.Request) string {
	hash := md5.Sum([]byte(r.Method + ":" + r.RequestURI))
	return fmt.Sprintf("%x", hash)
}

// Get retrieves a cached response if it exists and hasn't expired
func (rc *ResponseCache) Get(r *http.Request) (*CacheEntry, bool) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	key := rc.generateCacheKey(r)
	entry, exists := rc.cache[key]

	if !exists {
		return nil, false
	}

	// Check if entry has expired
	if time.Now().After(entry.ExpiresAt) {
		return nil, false
	}

	rc.log.Debug("cache hit", zap.String("key", key), zap.String("path", r.URL.Path))
	return entry, true
}

// Set stores a response in the cache
func (rc *ResponseCache) Set(r *http.Request, entry *CacheEntry) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	// Check cache size and evict oldest if necessary
	if len(rc.cache) >= rc.maxSize {
		rc.evictOldest()
	}

	key := rc.generateCacheKey(r)
	entry.ExpiresAt = time.Now().Add(rc.ttl)
	rc.cache[key] = entry

	rc.log.Debug("cache set", zap.String("key", key), zap.String("path", r.URL.Path))
}

// evictOldest removes the entry with the earliest expiration time
func (rc *ResponseCache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time = time.Now().Add(time.Hour * 24) // far future

	for key, entry := range rc.cache {
		if entry.ExpiresAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.ExpiresAt
		}
	}

	if oldestKey != "" {
		delete(rc.cache, oldestKey)
		rc.log.Debug("cache evicted", zap.String("key", oldestKey))
	}
}

// cleanupExpiredEntries removes expired entries periodically
func (rc *ResponseCache) cleanupExpiredEntries() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rc.mu.Lock()
		now := time.Now()
		count := 0

		for key, entry := range rc.cache {
			if now.After(entry.ExpiresAt) {
				delete(rc.cache, key)
				count++
			}
		}

		if count > 0 {
			rc.log.Debug("cleanup completed", zap.Int("entries_removed", count))
		}
		rc.mu.Unlock()
	}
}

// CachingMiddleware creates a middleware that caches responses for GET requests
func CachingMiddleware(cache *ResponseCache, log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Only cache GET requests
			if r.Method != http.MethodGet {
				next.ServeHTTP(w, r)
				return
			}

			// Check if response is cached
			if cached, ok := cache.Get(r); ok {
				// Write cached response
				for key, values := range cached.Headers {
					for _, value := range values {
						w.Header().Add(key, value)
					}
				}
				w.Header().Add("X-Cache", "HIT")
				w.WriteHeader(cached.StatusCode)

				// Inject cache status into cached response body
				body := string(cached.Body)
				if len(body) > 0 && body[len(body)-1] == '}' {
					body = body[:len(body)-1] + `,"cached":true}`
				}
				w.Write([]byte(body))
				return
			}

			// Wrap response writer to capture response for caching
			wrappedWriter := &cachedResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			// Wrap with cache status interceptor (for first-time hits)
			interceptor := &cacheInterceptorWriter{
				ResponseWriter: wrappedWriter,
				isCached:       false,
			}

			next.ServeHTTP(interceptor, r)

			// Cache successful GET responses (2xx status codes)
			if wrappedWriter.statusCode >= 200 && wrappedWriter.statusCode < 300 {
				entry := &CacheEntry{
					Body:       wrappedWriter.body,
					Headers:    wrappedWriter.Header(),
					StatusCode: wrappedWriter.statusCode,
				}
				cache.Set(r, entry)
			}
		})
	}
}

// cachedResponseWriter wraps http.ResponseWriter to capture response data
type cachedResponseWriter struct {
	http.ResponseWriter
	body       []byte
	statusCode int
	written    bool
}

func (rw *cachedResponseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.written = true
	}
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
}

func (rw *cachedResponseWriter) WriteHeader(statusCode int) {
	if !rw.written {
		rw.statusCode = statusCode
		rw.written = true
	}
	rw.ResponseWriter.WriteHeader(statusCode)
}

// cacheInterceptorWriter intercepts writes to inject cache status
type cacheInterceptorWriter struct {
	http.ResponseWriter
	isCached bool
	written  bool
}

func (w *cacheInterceptorWriter) Write(b []byte) (int, error) {
	if !w.written {
		// Inject cache status into the response body for first write
		body := string(b)
		if len(body) > 0 && body[len(body)-1] == '}' {
			body = body[:len(body)-1] + `,"cached":false}`
			b = []byte(body)
		}
		w.written = true
	}
	return w.ResponseWriter.Write(b)
}
