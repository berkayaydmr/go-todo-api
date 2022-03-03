package entities

import "time"

type ToDo struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	Details   string `gorm:"column:details;default:NewTo-Do "`
	Status    string `gorm:"column:status;default:To-Do"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
