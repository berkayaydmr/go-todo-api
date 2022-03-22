package mocks

import (
	"go-todo-api/entities"
	"go-todo-api/models"
	"time"
)

func User() *entities.User {
	Id := uint64(0)
	Email := "email"
	Password := ""
	Status := "Pending"
	return &entities.User{
		Id:       Id,
		Email:    Email,
		Password: Password,
		Status:   Status,
	}
}

func UserModelResponse() *models.UserResponse {
	Id := uint64(0)
	Email := "email"
	Status := "Pending"
	CreatedAt := time.Time{}
	UpdatedAt := time.Time{}
	return &models.UserResponse{
		Id:        Id,
		Email:     Email,
		Status:    Status,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}

func ToDo() *entities.ToDo {
	id := uint64(0)
	details := "testDetails"
	status := "On Progress"
	return &entities.ToDo{
		Id:        id,
		UserId:    uint64(0),
		User:      entities.User{Id: 0, Email: "", Status: "", Password: ""},
		Details:   details,
		Status:    status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func ToDoResponse() *models.ToDo {
	Id := uint64(0)
	Details := "testDetails"
	Status := "On Progress"
	CreatedAt := time.Time{}
	UpdatedAt := time.Time{}
	return &models.ToDo{
		ID:        Id,
		Details:   Details,
		Status:    Status,
		CreatedAt: CreatedAt.String(),
		UpdatedAt: UpdatedAt.String(),
	}
}

func ToDoPatchResponse() *models.ToDo {
	Id := uint64(0)
	Details := "Updated Detail"
	Status := "Done"
	CreatedAt := time.Time{}
	UpdatedAt := time.Time{}
	return &models.ToDo{
		ID:        Id,
		Details:   Details,
		Status:    Status,
		CreatedAt: CreatedAt.String(),
		UpdatedAt: UpdatedAt.String(),
	}
}
