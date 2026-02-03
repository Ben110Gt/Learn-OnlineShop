package models

import "time"

type Order struct {
	OrderID  string      `json:"order_id" gorm:"primaryKey"`
	UserID   string      `json:"user_id" gorm:"index"`
	User     User        `json:"user" gorm:"foreignKey:UserID;references:UserID"`
	Total    float64     `json:"total"`
	Status   string      `json:"status"`
	Items    []OrderItem `json:"items" gorm:"foreignKey:OrderID;references:OrderID"` // FK อยู่ที่ OrderItem.OrderID
	CreateAt time.Time   `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time   `json:"update_at" gorm:"autoUpdateTime"`
}
