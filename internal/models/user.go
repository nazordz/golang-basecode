package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string         `json:"name" gorm:"name;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Phone     string         `json:"phone" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	RoleID    string         `json:"role_id" gorm:"not null;"`
	Role      Role           `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserRequest struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Email    string `form:"email" binding:"required,email" json:"email"`
	Phone    string `form:"phone" binding:"required" json:"phone"`
	Password string `form:"password" binding:"required" json:"password"`
	Role     string `binding:"required,oneof=admin user guest" json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
