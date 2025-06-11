# REST CRUD Go API - A Simple Template

A clean, extensible and scalable REST API built with Go, featuring user management with CRUD operations, authentication, and pagination.

## 🚀 Features

- **Fast & Lightweight**: Built with Go's standard library and minimal dependencies.
- **Database Migrations**: Seamless schema management with Goose.
- **Database Integration**: PostgreSQL with connection pooling using pgx/v5.
- **Complete User Management**: Create, read, update, and delete users.
- **Authentication**: Secure basic login with good password hashing + HTTP-Only JWT Tokens.
- **Pagination**: Efficient data retrieval with limit/offset pagination.
- **Error Handling**: Proper HTTP status codes and error responses.
- **Input Validation**: Comprehensive request validation and sanitization.
- **JSON API**: RESTful endpoints with consistent response format.
- **OAuth2 Integration**: Secure authentication with GitHub OAuth2.

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Database Driver**: pgx/v5 with connection pooling
- **Database Migrations**: Goose
- **Password Hashing**: argon5
- **OAuth2 Handler**: goth

## 📁 Project Structure

```py
rest-crud-go/
├── internal/
│   ├── api/
│   │   ├── handlers/     # HTTP handlers
│   │   ├── middlewares/  # Middlewares
│   │   └── routes/       # Routes Endpoints
│   ├── core/
│   │   ├── services/     # Business logic layer
│   │   ├── repositories/ # Data access layer
│   │   └── models/       # Data models and DTOs
│   │
│   ├── database/ # Migrations
│   └── utils/    # Utility functions
│
├── tests/  # Simples Unit Tests
├── go.mod
├── go.sum
├── air.toml
├── README.md
└── main.go
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.23 or higher
- PostgreSQL 12+
- Git

### 1. Clone the repository

```bash
git clone https://github.com/Kazyel/crud-go.git
cd rest-crud-go
```

### 2. Install dependencies

```bash
go mod tidy
```

### 4. Configure environment variables

Create a `.env` file in the root directory:

```env
PORT=8080
DATABASE_URL=postgres_url

GITHUB_CLIENT_SECRET=secret_key
GITHUB_CLIENT_ID=client_id
GITHUB_CALLBACK_URL=callback_url
SESSION_SECRET=session_secret

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres_url
GOOSE_MIGRATION_DIR=./internal/database/migrations
GOOSE_TABLE=custom.goose_migrations

JWT_SECRET=jwt_secret
```

### 5. Run the application

```bash
go run main.go
```

The API will be available at `http://localhost:8080`

## Hot Reload Development

For development with automatic reloading, install Air:

```bash
go install github.com/cosmtrek/air@latest
```
