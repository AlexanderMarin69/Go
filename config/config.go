package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Config represents the application configuration
type Config struct {
	Server    ServerConfig   `json:"server"`
	RateLimit RateLimitConfig `json:"rateLimit"`
	Cache     CacheConfig    `json:"cache"`
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

// LoadConfig loads configuration from file
func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfgJSON struct {
		Server    ServerConfigJSON `json:"server"`
		RateLimit RateLimitConfig  `json:"rateLimit"`
		Cache     CacheConfigJSON  `json:"cache"`
	}

	if err := json.Unmarshal(data, &cfgJSON); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Parse durations and create final config
	cfg := &Config{
		RateLimit: cfgJSON.RateLimit,
	}

	// Parse server config
	cfg.Server = parseServerConfig(cfgJSON.Server)

	// Parse cache config
	cfg.Cache = parseCacheConfig(cfgJSON.Cache)

	return cfg, nil
}

// parseServerConfig converts string durations to time.Duration
func parseServerConfig(scJSON ServerConfigJSON) ServerConfig {
	sc := ServerConfig{
		Port:         scJSON.Port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if scJSON.Port == "" {
		sc.Port = ":8080"
	}

	if scJSON.ReadTimeout != "" {
		if d, err := time.ParseDuration(scJSON.ReadTimeout); err == nil {
			sc.ReadTimeout = d
		}
	}

	if scJSON.WriteTimeout != "" {
		if d, err := time.ParseDuration(scJSON.WriteTimeout); err == nil {
			sc.WriteTimeout = d
		}
	}

	if scJSON.IdleTimeout != "" {
		if d, err := time.ParseDuration(scJSON.IdleTimeout); err == nil {
			sc.IdleTimeout = d
		}
	}

	return sc
}

// parseCacheConfig converts string durations to time.Duration
func parseCacheConfig(ccJSON CacheConfigJSON) CacheConfig {
	cc := CacheConfig{
		Enabled: ccJSON.Enabled,
		TTL:     5 * time.Minute,
		MaxSize: 1000,
	}

	if ccJSON.TTL != "" {
		if d, err := time.ParseDuration(ccJSON.TTL); err == nil {
			cc.TTL = d
		}
	}

	if ccJSON.MaxSize > 0 {
		cc.MaxSize = ccJSON.MaxSize
	}

	return cc
}
