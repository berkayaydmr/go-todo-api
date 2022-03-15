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
		Id:        Id,
		Email:     Email,
		Password:  Password,
		Status:    Status,
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

func UserRequestModel() models.UserRequest {
	Email := "email"
	Password := "password"
	PasswordConfirm := "password"
	return models.UserRequest{
		Email:           Email,
		Password:        Password,
		PasswordConfirm: PasswordConfirm,
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

func ToDosResponse() []*entities.ToDo {
	return []*entities.ToDo{}
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
