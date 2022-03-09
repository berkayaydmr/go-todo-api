package mocks

import (
	"go-todo-api/entities"
	"time"
)

func ToDoResponse() *entities.ToDo {
	Id := uint64(0)
	Details := "testDetails"
	Status := "On Progress"
	CreatedAt := time.Time{}
	UpdatedAt := time.Time{}
	return &entities.ToDo{
		Id:        Id,
		Details:   Details,
		Status:    Status,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}

func ToDosResponse() []*entities.ToDo {
	return []*entities.ToDo{}
}


func ToDoPatchResponse() *entities.ToDo {
	Id := uint64(0)
	Details := "Updated Detail"
	Status := "Done"
	CreatedAt := time.Time{}
	UpdatedAt := time.Time{}
	return &entities.ToDo{
		Id:        Id,
		Details:   Details,
		Status:    Status,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}
