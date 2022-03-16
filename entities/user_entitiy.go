package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint64         `gorm:"primaryKey;column:id"`
	Email     string         `gorm:"column:email;"`
	Password  string         `gorm:"column:password"`
	Status    string         `gorm:"column:status;default:Pending"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}