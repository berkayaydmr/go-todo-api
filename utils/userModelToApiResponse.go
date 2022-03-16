package utils

import (
	"go-todo-api/entities"
	"go-todo-api/models"
)

func UserApiResponse(user *entities.User) (*models.UserResponse) {
	var userResponse = &models.UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return userResponse
}