package models

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID        string `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title     string
	Content   string
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
