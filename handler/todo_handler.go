package handler

import (
	"errors"
	"go-todo-api/entities"
	"go-todo-api/models"
	"go-todo-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ToDoHandler struct {
	ToDoRepository repository.ToDoRepositoryInterface
}

func NewToDoHandler(ToDoRepo repository.ToDoRepositoryInterface) *ToDoHandler {
	return &ToDoHandler{ToDoRepository: ToDoRepo}
}

func (handler *ToDoHandler) PostToDo(c *gin.Context) {
	logger,err := zap.NewProduction()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var createToDo = &models.ToDoRequest{}
	if err := c.BindJSON(&createToDo); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var newToDo = &entities.ToDo{}
	newToDo.Details = createToDo.Details
	newToDo.Status = createToDo.Status
	_, err = handler.ToDoRepository.Insert(newToDo)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusCreated, newToDo)
}

func (handler *ToDoHandler) PatchToDo(c *gin.Context) {
	logger,err := zap.NewProduction()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result.Id = id
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = handler.ToDoRepository.FindByID(&result)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	var patchToDo *models.ToDoRequest
	if err := c.ShouldBindJSON(&patchToDo); err != nil {
		logger.Error(err.Error())
		return
	}

	result.Status = patchToDo.Status
	result.Details = patchToDo.Details
	_, err = handler.ToDoRepository.Update(&result)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	logger,err := zap.NewProduction()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result.Id = id
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = handler.ToDoRepository.FindByID(&result)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	err = handler.ToDoRepository.Delete(&result)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (handler *ToDoHandler) GetToDos(c *gin.Context) {
	logger,err := zap.NewProduction()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var result = []*entities.ToDo{nil} //nil test için
	ToDos, err := handler.ToDoRepository.FindAll(result)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	//test için
	if len(ToDos) == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(200, ToDos)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	logger,err := zap.NewProduction()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result.Id = id
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = handler.ToDoRepository.FindByID(&result)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err.Error())
		c.JSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	c.JSON(200, result)
}
