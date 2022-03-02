package entities

import "time"

type ToDo struct {
	Id        uint64 `gorm:"primaryKey;column:id"`
	Details   string `gorm:"column:details"`
	Status    string `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
