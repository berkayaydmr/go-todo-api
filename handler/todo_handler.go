package handler

import (
	"go-todo-api/entities"
	"go-todo-api/models"
	"go-todo-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ToDoHandler struct {
	ToDoRepository repository.ToDoRepositoryInterface
}

func NewToDoHandler(ToDoRepo repository.ToDoRepositoryInterface) *ToDoHandler {
	return &ToDoHandler{ToDoRepository: ToDoRepo}
}

func (handler *ToDoHandler) PostToDo(c *gin.Context) {
	var createToDo = &models.ToDoRequest{}

	if err := c.BindJSON(&createToDo); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if createToDo.Validate() {
		zap.S().Error("Error: required field details send nil", createToDo.Details)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var newToDo = &entities.ToDo{
		Details: createToDo.Details,
		Status:  createToDo.Status,
		UserId:  userID,
	}

	err = handler.ToDoRepository.Insert(newToDo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var toDoResponse = &models.ToDo{
		ID: newToDo.Id,
		Details: newToDo.Details,
		Status: newToDo.Status,
		CreatedAt: newToDo.CreatedAt.String(),
		UpdatedAt: newToDo.UpdatedAt.String(),
	}

	c.JSON(http.StatusCreated, toDoResponse)
}

func (handler *ToDoHandler) PatchToDo(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: On User id convert", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	todoId, err := strconv.ParseUint(c.Param("todo_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: On Todo id convert", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var result = &entities.ToDo{
		Id:     todoId,
		UserId: userId,
	}

	todo, err := handler.ToDoRepository.FindByID(result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if todo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var patchToDo *models.ToDoPatchRequest
	if err := c.ShouldBindJSON(&patchToDo); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if patchToDo.Status != nil {
		todo.Status = *patchToDo.Status
	}
	if patchToDo.Details != nil {
		todo.Details = *patchToDo.Details
	}

	err = handler.ToDoRepository.Update(todo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var toDoResponse = &models.ToDo{
		ID: todo.Id,
		Details: todo.Details,
		Status: todo.Status,
		CreatedAt: todo.CreatedAt.String(),
		UpdatedAt: todo.UpdatedAt.String(),
	}
	c.JSON(http.StatusOK, toDoResponse)
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	todoId, err := strconv.ParseUint(c.Param("todo_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var result = &entities.ToDo{
		Id:     todoId,
		UserId: userId,
	}

	todo, err := handler.ToDoRepository.FindByID(result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if todo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	err = handler.ToDoRepository.Delete(todo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (handler *ToDoHandler) GetToDos(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var userID = &entities.ToDo{
		UserId: id,
	}

	toDos, err := handler.ToDoRepository.FindAll(userID)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var toDosResponse = make([]models.ToDo, len(toDos))

	for i := 0; i < len(toDos); i++ {
		toDosResponse[i] = models.ToDo{
			ID:        toDos[i].Id,
			UserId:    toDos[i].UserId,
			Details:   toDos[i].Details,
			Status:    toDos[i].Status,
			CreatedAt: toDos[i].CreatedAt.String(),
			UpdatedAt: toDos[i].UpdatedAt.String(),
		}
	}

	c.JSON(200, toDosResponse)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	toDoId, err := strconv.ParseUint(c.Param("todo_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var result = &entities.ToDo{
		Id:     toDoId,
		UserId: userId,
	}

	toDo, err := handler.ToDoRepository.FindByID(result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if toDo == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var toDoResponse = models.ToDo{
		ID:        toDo.Id,
		UserId:    toDo.UserId,
		Details:   toDo.Details,
		Status:    toDo.Status,
		CreatedAt: toDo.CreatedAt.String(),
		UpdatedAt: toDo.UpdatedAt.String(),
	}

	c.JSON(200, toDoResponse)
}
