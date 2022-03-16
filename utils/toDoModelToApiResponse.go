package utils

import (
	"go-todo-api/entities"
	"go-todo-api/models"
)

func ToDoApiResponse(toDo *entities.ToDo) (*models.ToDo) {
	var toDoResponse = &models.ToDo{
		ID:        toDo.Id,
		UserId:    toDo.UserId,
		Details:   toDo.Details,
		Status:    toDo.Status,
		CreatedAt: toDo.CreatedAt.String(),
		UpdatedAt: toDo.UpdatedAt.String(),
	}
	return toDoResponse
}