package handler

import (
	"fmt"
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
	fmt.Println(createToDo.Details)
	if createToDo.Validate() {
		zap.S().Error("Error: required field details send nil", createToDo.Details)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var newToDo = &entities.ToDo{
		Details: createToDo.Details,
		Status:  createToDo.Status,
	}
	err := handler.ToDoRepository.Insert(newToDo)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusCreated, newToDo)
}

func (handler *ToDoHandler) PatchToDo(c *gin.Context) {
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result.Id = id
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	record, err := handler.ToDoRepository.FindByID(&result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if record == nil {
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
		result.Status = *patchToDo.Status
	}
	if patchToDo.Details != nil {
		result.Details = *patchToDo.Details
	}

	_, err = handler.ToDoRepository.Update(&result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result.Id = id
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	record, err := handler.ToDoRepository.FindByID(&result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if record == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	err = handler.ToDoRepository.Delete(&result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (handler *ToDoHandler) GetToDos(c *gin.Context) {
	var result []*entities.ToDo
	ToDos, err := handler.ToDoRepository.FindAll(result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(200, ToDos)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result.Id = id
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	record, err := handler.ToDoRepository.FindByID(&result)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if record == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(200, result)
}
