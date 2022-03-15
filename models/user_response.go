package models

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        uint64         `json:"id"`
	Email     string         `json:"email"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
