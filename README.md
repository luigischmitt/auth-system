# 🔐 Auth System - JWT Authentication API

A complete authentication system built with Go, featuring JWT tokens, MongoDB, and RESTful API endpoints. This project implements secure user registration, login, and protected routes using industry-standard practices.

## Features

- 🔑 **JWT Authentication** - Secure token-based authentication
- 🔒 **Password Hashing** - bcrypt for secure password storage
- 🍪 **Cookie-based Sessions** - HTTP-only cookies for token storage
- 🛡️ **Protected Routes** - Middleware-based route protection
- 📊 **MongoDB Integration** - NoSQL database with MongoDB driver v2
- 🏗️ **Clean Architecture** - Separation of concerns with repository pattern
- 🐳 **Docker Support** - Containerized MongoDB setup
- ⚡ **High Performance** - Built with Gin web framework
- 🔧 **Environment Configuration** - Flexible environment-based config

## Quick Start

### Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/auth-system.git
   cd auth-system
   ```

2. **Set up environment variables**
   ```bash
   cp env.example .env
   # Edit .env with your configurations
   ```

3. **Start MongoDB with Docker**
   ```bash
   docker compose up -d
   ```

4. **Install dependencies**
   ```bash
   go mod download
   ```

5. **Run the application**
   ```bash
   go run ./cmd/api
   ```

The server will start on `http://localhost:8080`

## 📁 Project Structure

```
auth-system/
├── cmd/
│   └── api/                 # Application entrypoint
│       └── main.go
├── internal/
│   ├── auth/               # Authentication utilities
│   │   ├── password.go     # Password hashing & validation
│   │   └── tokens.go       # JWT token management
│   ├── database/           # Database connection
│   │   └── database.go     # MongoDB service
│   ├── middlewares/        # HTTP middlewares
│   │   └── auth.go         # JWT authentication middleware
│   ├── models/             # Data models
│   │   └── user.go         # User model & validation
│   ├── repositories/       # Data access layer
│   │   └── user-repositories.go
│   └── server/             # HTTP server & handlers
│       ├── server.go       # Server setup & routing
│       ├── auth-handlers.go # Authentication endpoints
│       └── user-handlers.go # User endpoints
├── compose.yml             # Docker Compose configuration
├── go.mod                  # Go module definition
└── .env                    # Environment variables
```

## 🔌 API Endpoints

### Authentication Routes

| Method | Endpoint | Description | Body |
|--------|----------|-------------|------|
| `POST` | `/api/auth/register` | Register new user | `{"username", "email", "password"}` |
| `POST` | `/api/auth/login` | User login | `{"email", "password"}` |
| `POST` | `/api/auth/logout` | User logout | - |
| `POST` | `/api/auth/refresh` | Refresh access token | - |

### Protected Routes

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/api/user` | Get current user profile | ✅ |

## Testing the API

### Register a new user
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com", 
    "password": "securepassword123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

### Access protected route
```bash
curl -X GET http://localhost:8080/api/user \
  -H "Content-Type: application/json" \
  --cookie-jar cookies.txt --cookie cookies.txt
```

## ⚙️ Environment Variables

Create a `.env` file in the root directory:

```env
MONGODB_URL="mongodb://admin:password123@localhost:27017/auth-system?authSource=admin"
MONGODB_DATABASE=auth-system
ACCESS_TOKEN_SECRET=your-super-secret-jwt-key-here
PORT=8080
```

## Architecture

### Key Components

- **Repository Pattern**: Abstraction over data access
- **Dependency Injection**: Loose coupling between layers
- **Middleware Chain**: Authentication and request processing
- **Environment Configuration**: Flexible deployment options

## 🔒 Security Features

- **Password Hashing**: bcrypt with salt rounds
- **JWT Tokens**: HS256 algorithm with configurable expiration
- **HTTP-Only Cookies**: XSS protection
- **Input Validation**: Request body validation
- **Error Handling**: Secure error responses

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
