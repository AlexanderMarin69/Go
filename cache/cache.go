package cache

import (
	"encoding/json"
	"fmt"
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

// NewService creates a new cache service
func NewService(ttl time.Duration, maxSize int) *Service {
	cs := &Service{
		storage: make(map[string]*Entry),
		ttl:     ttl,
		maxSize: maxSize,
	}

	// Start background cleanup
	go cs.cleanupExpired()

	return cs
}

// Get retrieves a value from cache if it exists and hasn't expired
func (cs *Service) Get(key string) (interface{}, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	entry, exists := cs.storage[key]
	if !exists {
		return nil, false
	}

	// Check if expired
	if time.Now().After(entry.ExpiresAt) {
		return nil, false
	}

	return entry.Value, true
}

// GetJSON retrieves and unmarshals a cached value into the provided struct
func (cs *Service) GetJSON(key string, v interface{}) bool {
	val, exists := cs.Get(key)
	if !exists {
		return false
	}

	// If value is a string (JSON), unmarshal it
	if jsonStr, ok := val.(string); ok {
		return json.Unmarshal([]byte(jsonStr), v) == nil
	}

	// If value is already the expected type
	if jsonVal, ok := val.(interface{}); ok {
		if b, err := json.Marshal(jsonVal); err == nil {
			return json.Unmarshal(b, v) == nil
		}
	}

	return false
}

// Set stores a value in cache with TTL
func (cs *Service) Set(key string, value interface{}) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	// Check if cache is at max size, evict oldest if needed
	if len(cs.storage) >= cs.maxSize && cs.storage[key] == nil {
		cs.evictOldest()
	}

	cs.storage[key] = &Entry{
		Value:     value,
		ExpiresAt: time.Now().Add(cs.ttl),
	}
}

// SetJSON stores a JSON-serializable value in cache
func (cs *Service) SetJSON(key string, value interface{}) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	cs.Set(key, string(jsonBytes))
	return nil
}

// Contains checks if a key exists in cache and hasn't expired
func (cs *Service) Contains(key string) bool {
	_, exists := cs.Get(key)
	return exists
}

// Remove deletes a key from cache
func (cs *Service) Remove(key string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	delete(cs.storage, key)
}

// Clear removes all entries from cache
func (cs *Service) Clear() {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.storage = make(map[string]*Entry)
}

// evictOldest removes the entry with the earliest expiration time
func (cs *Service) evictOldest() {
	var oldestKey string
	var oldestTime time.Time = time.Now().Add(time.Hour * 24) // far future

	for key, entry := range cs.storage {
		if entry.ExpiresAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.ExpiresAt
		}
	}

	if oldestKey != "" {
		delete(cs.storage, oldestKey)
	}
}

// cleanupExpired removes expired entries periodically
func (cs *Service) cleanupExpired() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		cs.mu.Lock()
		now := time.Now()

		for key, entry := range cs.storage {
			if now.After(entry.ExpiresAt) {
				delete(cs.storage, key)
			}
		}

		cs.mu.Unlock()
	}
}

// Size returns the number of items currently in cache
func (cs *Service) Size() int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	return len(cs.storage)
}
