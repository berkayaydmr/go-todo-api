package entities

import "time"

type ToDo struct {
	Id        uint64    `gorm:"primaryKey;column:todo_id"`
	UserId    uint64    `gorm:"column:user_id"`
	User      User      `gorm:"foreingKey:UserId;"`
	Details   string    `gorm:"column:details;"`
	Status    string    `gorm:"column:status;default:To-Do"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
