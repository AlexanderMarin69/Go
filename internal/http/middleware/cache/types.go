package cache

import (
	"sync"
	"time"
)

// Entry represents a cached item with expiration
type Entry struct {
	Value     interface{}
	ExpiresAt time.Time
}

// Service provides caching operations with TTL support
type Service struct {
	mu      sync.RWMutex
	storage map[string]*Entry
	ttl     time.Duration
	maxSize int
}
