package models

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        uint64         `json:"Id"`
	Email     string         `json:"Email"`
	Status    string         `json:"Status"`
	CreatedAt time.Time      `json:"Created_at"`
	UpdatedAt time.Time      `json:"Updated_at"`
	DeletedAt gorm.DeletedAt `json:"Deleted_at"`
}
