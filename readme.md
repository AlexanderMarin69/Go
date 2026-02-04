# Go Web App - Production-Ready API Server

A modern, idiomatic Go web application built with production standards, concurrent architecture, and RESTful API design. Structured for scalability and maintainability with clear separation of concerns.

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
├── go.mod                   # Module dependencies
├── logger/
│   └── logger.go           # Structured logging setup (Zap)
├── middleware/
│   └── logging.go          # HTTP request logging middleware
├── controllers/
│   ├── user_controller.go      # User API v1 endpoints
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

### 2. Run the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

You should see output similar to:
```
{"level":"info","ts":1707000000.000000,"caller":"main.go:22","msg":"starting web server"}
{"level":"info","ts":1707000000.000001,"caller":"main.go:33","msg":"listening on address","addr":":8080"}
```

### 3. Health Check

Verify the server is running:
```bash
curl http://localhost:8080/health
```

Expected response: `{"status":"ok"}`

### 4. Graceful Shutdown

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

### 🔄 Concurrency & Performance
- Goroutine-based server startup and request handling
- Non-blocking HTTP server with timeouts
- Efficient request logging without blocking

### 📋 Code Quality
- Production-ready error handling
- Helper functions for route registration (DRY principle)
- Idiomatic Go naming and patterns
- Easy onboarding for new developers

## Dependencies

- `github.com/go-chi/chi/v5` - HTTP router and middleware framework
- `go.uber.org/zap` - Structured logging library

## Next Steps (Future Sections)

- Section 2: Rate Limiting
- Section 3: Caching
- Section 4: Authentication & Authorization
- Section 5: Database / ORM
- ... and more

---

**Built with Go best practices and production standards in mind.**
