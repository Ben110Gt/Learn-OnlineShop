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

# Project Structure
 
```
Learn-OnlineShop/
├── cmd/
│   └── api/
│       └── main.go                  # Application entry point
├── internal/
│   ├── configs/
│   │   └── DB.go                    # Database connection & AutoMigrate
│   ├── Handler/                     # HTTP handlers (controllers)
│   │   ├── cart_handler.go
│   │   ├── category_handler.go
│   │   ├── order_handler.go
│   │   ├── product_handler.go
│   │   └── user_handler.go
│   ├── middlewares/
│   │   └── middleware.go            # JWT auth, AdminOnly, UserOnly guards
│   ├── models/                      # GORM data models
│   │   ├── Auth.go
│   │   ├── Cart.go
│   │   ├── Category.go
│   │   ├── Oder.go
│   │   ├── OderItem.go
│   │   ├── Product.go
│   │   └── User.go
│   ├── repositories/                # Database access layer
│   │   ├── cart_repositories.go
│   │   ├── category_repository.go
│   │   ├── order_repositories.go
│   │   ├── product_repositories.go
│   │   └── user_repositories.go
│   ├── Routes/                      # Route definitions
│   │   ├── cart_routes.go
│   │   ├── category_routes.go
│   │   ├── order_routes.go
│   │   ├── product_routes.go
│   │   ├── public.go
│   │   └── user_routes.go
│   ├── server/
│   │   └── server.go                # Fiber app setup & CORS config
│   ├── Service/                     # Business logic layer
│   │   ├── cart_service.go
│   │   ├── category_service.go
│   │   ├── order_service.go
│   │   ├── product_service.go
│   │   └── user_service.go
│   ├── Upload/                      # Static file storage
│   │   ├── Product/
│   │   └── User/
│   └── util/                        # Utility helpers
│       ├── Check_Passworld.go
│       ├── Generate_ID.go
│       ├── Hash_Password.go
│       ├── upload_product_Bill.go
│       └── upload_User_File.go
├── .env
├── go.mod
└── go.sum
```
 
---
 
## Prerequisites
 
- **Go** `1.25.0` or later
- **PostgreSQL** running locally or on a reachable host
Verify your Go installation:
 
```bash
go version
```
 
---
 
## Environment Configuration
 
Create a `.env` file in the project root (same directory as `go.mod`) with the following variables:
 
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=E-Commerce
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Bangkok
ACCESS_TOKEN_SECRET=your_jwt_secret
```
 
> These variable names match what is read in `internal/configs/DB.go`. Adjust values to match your local PostgreSQL setup.
 
---
 
## Getting Started
 
**1. Clone / extract the project**
 
```bash
# or extract the zip, then navigate into the folder
cd Learn-OnlineShop
```
 
**2. Install dependencies**
 
```bash
go mod tidy
```
 
**3. Configure the database**
 
Make sure PostgreSQL is running and the database specified in `DB_NAME` exists. GORM will run `AutoMigrate` on startup and create all required tables automatically.
 
**4. Run the server**
 
```bash
go run ./cmd/api
```
 
The server starts at `http://localhost:8080`.
 
**5. Build a binary (optional)**
 
```bash
go build -o onlineshop-api ./cmd/api
./onlineshop-api
```
 
On Windows:
 
```powershell
go build -o onlineshop-api.exe ./cmd/api
.\onlineshop-api.exe
```
 
---
 
## API Reference
 
### Authentication
 
Protected routes require a JWT token in the `Authorization` header:
 
```
Authorization: Bearer <token>
```
 
Tokens are issued at login. Two roles exist: **`user`** and **`admin`**.
 
---
 
### Public Endpoints
 
No authentication required.
 
| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/register` | Register a new user account |
| `POST` | `/login` | Log in and receive a JWT token |
| `GET` | `/products` | List all published products |
| `GET` | `/categories` | List all product categories |
 
---
 
### User Endpoints
 
Requires `Authorization: Bearer <token>` with role **`user`**.
 
#### Cart
 
| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/user/cart` | Add an item to the cart |
| `GET` | `/user/cart/:user_id` | Get cart contents for a user |
 
#### Order
 
| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/user/order` | Create an order from cart items |
 
---
 
### Admin Endpoints
 
Requires `Authorization: Bearer <token>` with role **`admin`**.
 
#### Users
 
| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/admin/users` | List all users |
| `GET` | `/admin/user/:user_id` | Get a single user by ID |
| `POST` | `/admin/register/` | Register a new admin account |
| `PUT` | `/admin/user/:user_id` | Update a user |
| `DELETE` | `/admin/user/:user_id` | Delete a user |
 
#### Products
 
| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/admin/products` | List all products (admin view) |
| `GET` | `/admin/product/:product_id` | Get a product by ID |
| `POST` | `/admin/products` | Create a new product |
| `PUT` | `/admin/product/:product_id` | Update a product |
| `DELETE` | `/admin/product/:product_id` | Delete a product |
 
#### Categories
 
| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/admin/categories` | List all categories |
| `POST` | `/admin/categories` | Create a new category |
| `DELETE` | `/admin/categories/:category_id` | Delete a category |
 
#### Orders
 
| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/admin/orders/` | List all orders |
| `POST` | `/admin/order/:order_id` | Confirm / update an order status |
 
---

## Author
 
- Name: Souksakhone Haknolath
- Contact: souksakhone.haknolath@gmail.com


