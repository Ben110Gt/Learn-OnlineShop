package models

import "time"

type User struct {
	UserID       string    `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Password     string    `json:"password" gorm:"not null"`
	Phone        string    `json:"phone"`
	Role         string    `json:"role"`
	ProfileImage string    `json:"profile_image"`
	CreateAt     time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt     time.Time `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt     time.Time `json:"delete_at" gorm:"index"`
}
