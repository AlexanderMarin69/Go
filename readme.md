# Go Web App - .NET like DIY API Server

Learning about Go server code coming from .NET.
**Trying to** make modern, idiomatic Go web application built with production standards, concurrent architecture, and RESTful API design. Structured for scalability and maintainability with clear separation of concerns.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Features](#features)

## Prerequisites

- **Go 1.25.6 or higher** - [Download Go](https://golang.org/dl/)
- **Postman** (optional, for API testing) - [Download Postman](https://www.postman.com/downloads/)

## Project Structure

```
├── main.go                  # Server entry point, graceful shutdown
├── config.json              # Application configuration (server, rate limit, cache)
├── go.mod                   # Module dependencies
├── cache/
│   └── cache.go            # Cache service with DI (Get, Set, Contains, Remove, Clear)
├── config/
│   └── config.go           # Configuration loading and defaults
├── logger/
│   └── logger.go           # Structured logging setup (Zap)
├── middleware/
│   ├── logging.go          # HTTP request logging middleware
│   └── rate_limit.go       # Rate limiting middleware with token bucket
├── controllers/
│   ├── user_controller.go      # User API v1 endpoints with cache logic
│   ├── response_helpers.go     # Shared response helper functions
│   ├── user_v2_controller.go   # User API v2 endpoints
│   └── product_controller.go   # Product API v1 endpoints
├── routes/
│   └── routes.go           # Router configuration and route registration
└── readme.md               # This file
```

## Getting Started

### 1. Install Dependencies

```bash
go mod download
```

### 2. Configure the Application (Optional)

Edit `config.json` to customize server settings:

```json
{
  "server": {
    "port": ":8080",
    "readTimeout": "15s",
    "writeTimeout": "15s",
    "idleTimeout": "60s"
  },
  "rateLimit": {
    "enabled": true,
    "requestsPerSec": 100,
    "burstSize": 10
  },
  "cache": {
    "enabled": true,
    "ttl": "5m",
    "maxSize": 1000
  }
}
```

### 3. Run the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080` (or custom port from config)

You should see output similar to:
```
{"level":"info","ts":1707000000.000000,"caller":"main.go:22","msg":"starting web server"}
{"level":"info","ts":1707000000.000001,"caller":"main.go:42","msg":"listening on address","addr":":8080"}
```

### 4. Health Check

Verify the server is running:
```bash
curl http://localhost:8080/health
```

Expected response: `{"status":"ok"}`

### 5. Graceful Shutdown

Press `Ctrl+C` to gracefully shut down the server. The server will complete in-flight requests within 5 seconds.

## API Endpoints

### Health Check
- **GET** `/health` - Server health status

### API v1 - User Endpoints
- **GET** `/api/v1/users` - List all users
- **POST** `/api/v1/users` - Create new user (Returns 201)
- **GET** `/api/v1/users/{id}` - Get user by ID
- **PUT** `/api/v1/users/{id}` - Update user by ID
- **DELETE** `/api/v1/users/{id}` - Delete user (Returns 400 for testing)

### API v1 - Product Endpoints
- **GET** `/api/v1/products` - List all products
- **POST** `/api/v1/products` - Create new product (Returns 201)
- **GET** `/api/v1/products/{id}` - Get product by ID
- **PUT** `/api/v1/products/{id}` - Update product by ID
- **DELETE** `/api/v1/products/{id}` - Delete product (Returns 400 for testing)

### API v2 - User Endpoints (Enhanced Format)
- **GET** `/api/v2/users` - List all users (v2 format with pagination)
- **GET** `/api/v2/users/{id}` - Get user by ID (v2 format)

## Testing with Postman

### Quick Setup

1. Open Postman
(1.5 - import postman collection from this github repo)
2. Create a new request
3. Use the following examples:

#### Example 1: GET - List Users (v1)
```
GET http://localhost:8080/api/v1/users
```

#### Example 2: POST - Create User (v1)
```
POST http://localhost:8080/api/v1/users
Content-Type: application/json

Body: (any JSON)
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

#### Example 3: GET - User by ID (v1)
```
GET http://localhost:8080/api/v1/users/123
```

#### Example 4: DELETE - Delete User (Returns 400)
```
DELETE http://localhost:8080/api/v1/users/123
```

#### Example 5: Compare Versions
Compare responses from:
- `GET http://localhost:8080/api/v1/users` (v1 format)
- `GET http://localhost:8080/api/v2/users` (v2 format with pagination)

## Features

### ✅ Section 1: Web Framework / Router
- **Chi Router** - Lightweight, fast, idiomatic Go router
- **Structured Logging** - Production-ready structured logging with Zap
- **API Versioning** - Support for multiple API versions (`/api/v1`, `/api/v2`)
- **Controller Pattern** - .NET-style controller organization (multiple files per domain)
- **Middleware Architecture** - Extensible middleware system (logging implemented)
- **Graceful Shutdown** - Proper signal handling and server shutdown
- **Error Handling** - Structured error responses with status codes
- **Response Helpers** - Reusable helper functions for consistent response formatting

### ✅ Section 2: Rate Limiting
- **Token Bucket Algorithm** - Configurable requests per second and burst size
- **Configurable** - Enable/disable and adjust limits via `config.json`
- **429 Status Responses** - Proper HTTP status codes for rate limit exceeded
- **Structured Logging** - Logs rate limit violations with IP and path information
- **Concurrent Safe** - Thread-safe rate limiting using Go's token bucket

### ✅ Section 3: Response Caching
- **Dependency Injection** - Cache service injected into controllers for explicit control
- **Manual Cache Logic** - Handlers control exactly when/what to cache (like .NET)
- **TTL Support** - Configurable cache time-to-live (default: 5 minutes)
- **LRU Eviction** - Oldest entries removed when cache reaches max size
- **Background Cleanup** - Periodic cleanup of expired cache entries (every 1 minute)
- **Thread-Safe** - Concurrent-safe caching using sync.RWMutex
- **Rich API** - Get(), Set(), Contains(), Remove(), Clear() methods
- **JSON Support** - GetJSON() and SetJSON() for automatic serialization
- **Configurable** - Enable/disable and adjust TTL/size via `config.json`
- **Cache Status** - Response includes `"cached": true/false` field

### Example: Using Cache in Handlers

```go
// In your handler with cache service injected
if cacheService != nil {
    // Try to get from cache
    if cached, ok := cacheService.Get("users:list"); ok {
        writeSuccessWithCacheStatus(w, http.StatusOK, cached.(string), true)
        return
    }
}

// Not in cache, generate response
message := "list all users"

// Store in cache for next time
if cacheService != nil {
    cacheService.Set("users:list", message)
}

writeSuccessWithCacheStatus(w, http.StatusOK, message, false)
```

## Configuration

### config.json Structure

- **server.port** - Server listen address (default: `:8080`)
- **server.readTimeout** - HTTP read timeout (default: `15s`)
- **server.writeTimeout** - HTTP write timeout (default: `15s`)
- **server.idleTimeout** - HTTP idle timeout (default: `60s`)
- **rateLimit.enabled** - Enable/disable rate limiting (default: `true`)
- **rateLimit.requestsPerSec** - Allowed requests per second (default: `100`)
- **rateLimit.burstSize** - Burst capacity for token bucket (default: `10`)
- **cache.enabled** - Enable/disable response caching (default: `true`)
- **cache.ttl** - Cache entry time-to-live (default: `5m`)
- **cache.maxSize** - Maximum number of cached entries (default: `1000`)

## Dependencies

- `github.com/go-chi/chi/v5` - HTTP router and middleware framework
- `go.uber.org/zap` - Structured logging library
- `golang.org/x/time` - Token bucket rate limiting

## Next Steps (Future Sections)

- Section 4: Authentication & Authorization
- Section 5: Database / ORM
- ... and more



