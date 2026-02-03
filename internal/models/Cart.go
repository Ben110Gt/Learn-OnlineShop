package models

import "time"

type Cart struct {
	CartID    string    `gorm:"primaryKey" json:"cart_id"`
	UserID    string    `json:"user_id" gorm:"index"`
	ProductID string    `json:"product_id" gorm:"index"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID;references:ProductID"`
	Quantity  int       `json:"quantity"`
	CreateAt  time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt  time.Time `json:"update_at" gorm:"autoUpdateTime"`
}
