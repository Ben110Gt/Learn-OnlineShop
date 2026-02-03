package models

import "time"

type Category struct {
	CategoryID   string    `json:"category_id" gorm:"primaryKey"`
	CategoryName string    `json:"categoryname"`
	CreateAt     time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt     time.Time `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt     time.Time `json:"delete_at" gorm:"index"`
}
