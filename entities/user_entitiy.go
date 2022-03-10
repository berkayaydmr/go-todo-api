package entities

import "time"

type User struct {
	Id        uint64    `gorm:"primaryKey;column:user_id"`
	Email     string    `gorm:";column:email;"`
	Password  string    `gorm:"column:pasword"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
