# Go Web App - .NET like DIY API Server

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Features](#features)

## Project Structure

```
├── main.go                  # Server entry point, graceful shutdown
├── config.json              # Application configuration (server, rate limit, cache)
├── go.mod                   # Module dependencies
├── cache/
│   ├── types.go            # Cache service types (Entry, Service structs)
│   └── cache.go            # Cache service implementation (Get, Set, Contains, Remove, Clear)
├── config/
│   ├── types.go            # Configuration types (Config, ServerConfig, RateLimit, Cache structs)
│   └── config.go           # Configuration loading and parsing logic
├── logger/
│   └── logger.go           # Structured logging setup (Zap)
├── middleware/
│   ├── types.go            # Middleware types (responseWriter struct)
│   ├── logging.go          # HTTP request logging middleware
│   └── rate_limit.go       # Rate limiting middleware with token bucket
├── controllers/
│   ├── user_controller.go      # User API v1 endpoints with cache logic
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

Press `Ctrl+C` to shut down the server. 
## API Endpoints

### Health Check
- **GET** `/health` - Server health status

### API v1 - User Endpoints
- **GET** `/api/v1/users` - List all users
- **POST** `/api/v1/users` - Create new user (Returns 201)
- **GET** `/api/v1/users/{id}` - Get user by ID
- **PUT** `/api/v1/users/{id}` - Update user by ID
- **DELETE** `/api/v1/users/{id}` - Delete user (Returns 400 for testing)

## Testing with Postman

### Quick Setup

1. Open Postman
(1.5 - import postman collection from this github repo)
2. Create a new request
3. Use the following examples:

## Roadmap

