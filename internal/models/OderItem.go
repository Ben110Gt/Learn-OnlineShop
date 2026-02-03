package models

import "time"

type OrderItem struct {
	OrderItemID string  `json:"order_item_id" gorm:"primaryKey"`
	OrderID     string  `json:"order_id" gorm:"index"`
	ProductID   string  `json:"product_id" gorm:"index"`
	Product     Product `json:"product" gorm:"foreignKey:ProductID;references:ProductID"`
	Quantity    float64
	Price       float64
	CreateAt    time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoUpdateTime"`
}
