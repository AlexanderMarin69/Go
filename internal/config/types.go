package config

import "time"

// Config represents the application configuration
type Config struct {
	Server    ServerConfig    `json:"server"`
	RateLimit RateLimitConfig `json:"rateLimit"`
	Cache     CacheConfig     `json:"cache"`
}

// ServerConfigJSON is the JSON representation of server config with string durations
type ServerConfigJSON struct {
	Port         string `json:"port"`
	ReadTimeout  string `json:"readTimeout"`
	WriteTimeout string `json:"writeTimeout"`
	IdleTimeout  string `json:"idleTimeout"`
}

// ServerConfig contains HTTP server configuration with parsed durations
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// RateLimitConfig contains rate limiting configuration
type RateLimitConfig struct {
	Enabled        bool `json:"enabled"`
	RequestsPerSec int  `json:"requestsPerSec"`
	BurstSize      int  `json:"burstSize"`
}

// CacheConfigJSON is the JSON representation of cache config with string durations
type CacheConfigJSON struct {
	Enabled bool   `json:"enabled"`
	TTL     string `json:"ttl"`
	MaxSize int    `json:"maxSize"`
}

// CacheConfig contains response caching configuration with parsed durations
type CacheConfig struct {
	Enabled bool
	TTL     time.Duration
	MaxSize int
}
