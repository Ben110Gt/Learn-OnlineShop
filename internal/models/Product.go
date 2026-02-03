package models

import "time"

type Product struct {
	ProductID    string    `json:"product_id" gorm:"primaryKey"`
	ProductName  string    `json:"productname" gorm:"not null"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	CategoryID   string    `json:"category_id"`
	Category     Category  `json:"category"  gorm:"foreignKey:CategoryID;references:CategoryID"`
	ProductImage string    `json:"image"`
	Stock        int       `json:"stok"`
	CreateAt     time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt     time.Time `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt     time.Time `json:"delete_at" gorm:"index"`
}
