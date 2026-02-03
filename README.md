## Online Shop API (Go + Fiber + GORM)

A RESTful API for an online shop system built with **Go**, using **Fiber** as the HTTP framework and **GORM** for working with a **PostgreSQL** database.  
It provides core modules such as Users, Products, Categories, Cart, and Orders.

---

## Tech Stack

- **Language**: Go `go1.25.0`
- **Framework**: `github.com/gofiber/fiber/v2`
- **ORM**: `gorm.io/gorm` + `gorm.io/driver/postgres`
- **Auth**: `github.com/golang-jwt/jwt/v5`
- **Environment variables**: `github.com/joho/godotenv`

---

## Project Structure (overview)

```text
cmd/
  api/
    main.go              # Application entrypoint
internal/
  configs/
    DB.go                # Database connection and configuration
  Handler/               # HTTP handlers (controllers) for each resource
  middlewares/           # Middlewares (e.g. auth, logging)
  models/                # Data models (database schema)
  repositories/          # Data access layer (repositories)
  Routes/                # API routes mapping
  server/
    server.go            # Fiber server setup
  Service/               # Business logic (service layer)
  Upload/                # Example image files for products/users
  util/                  # Utility functions (hashing, file upload, etc.)
go.mod
go.sum
```

---

## Prerequisites

- **Go** installed (compatible with version specified in `go.mod`)
- **PostgreSQL** installed and running

Check Go installation:

```bash
go version
```

---


## Install Dependencies

From the project root, run:

```bash
go mod tidy
```

This will download and tidy all dependencies defined in `go.mod`.

---

## Run the Server

Run the API with:

```bash
go run ./cmd/api
```

By default, the server should listen on `http://localhost:<APP_PORT>` (e.g. `http://localhost:8080`),  
depending on how `APP_PORT` or the port in `server/server.go` is configured.

---

## Main Features (overview)

- **User**
  - Register / login / manage user profile
  - Passwords stored as secure hashes (see `internal/util/Hash_Password.go`)
- **Product**
  - Create, update, delete, and list products
  - Support for product image upload (see `internal/Upload/Product` and related util functions)
- **Category**
  - Manage product categories
- **Cart**
  - Add/remove items in a user’s cart
  - View cart details
- **Order**
  - Create orders from cart items
  - Store Order and OrderItem records

For more details on each feature:

- `internal/models/`       – data models / database schema
- `internal/Handler/`      – HTTP handlers (controllers)
- `internal/Routes/`       – route definitions mapping URLs → handlers
- `internal/Service/`      – business logic
- `internal/repositories/` – database access logic

---


