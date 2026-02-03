package models

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
type Claims struct {
	UserID string `json:"user_id"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
type ProductResponse struct {
	ProductID     string  `json:"product_id"`
	ProductName   string  `json:"productname"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CategoryID    string  `json:"category_id"`
	Category      string  `json:"category"`
	Product_Image string  `json:"image"`
	Stock         int     `json:"stok"`
}
type CartResponse struct {
	CartID      string `json:"cart_id"`
	UserID      string `json:"user_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	// Price       int    `json:"price"`
	Quantity int `json:"quantity"`
}
type CreateOrderRequest struct {
	UserID string `json:"user_id"`
}
type OrderItemResponse struct {
	OrderItemID string `json:"order_item_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type OrderResponse struct {
	OrderID string              `json:"order_id"`
	UserID  string              `json:"user_id"`
	Total   int                 `json:"total"`
	Status  string              `json:"status"`
	Items   []OrderItemResponse `json:"items"`
}
type ProductResponses struct {
	ProductID     string  `json:"product_id"`
	ProductName   string  `json:"productname"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CategoryID    string  `json:"category_id"`
	Category      string  `json:"category"`
	Product_Image string  `json:"image"`
}
