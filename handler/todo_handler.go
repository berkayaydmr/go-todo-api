package handler

import (
	"fmt"
	"go-todo-api/entities"
	"go-todo-api/models"
	"go-todo-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ToDoHandler struct {
	ToDoRepository repository.ToDoRepositoryInterface
}

func NewToDoHandler(ToDoRepo repository.ToDoRepositoryInterface) *ToDoHandler {
	return &ToDoHandler{ToDoRepository: ToDoRepo}
}

func (handler *ToDoHandler) PostToDo(c *gin.Context) {
	var CreateToDo = &models.ToDoRequest{}

	if err := c.BindJSON(&CreateToDo); err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var NewToDo = &entities.ToDo{}
	NewToDo.Details = CreateToDo.Details
	NewToDo.Status = CreateToDo.Status
	handler.ToDoRepository.Insert(NewToDo)
	c.JSON(http.StatusCreated, NewToDo)
}

func (handler *ToDoHandler) PatchToDo(c *gin.Context) {
	var result entities.ToDo
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = handler.ToDoRepository.FindByID(&result, int(id))
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusNotFound, nil)
		return
	}
	var PatchToDo *models.ToDoRequest
	if err := c.ShouldBindJSON(&PatchToDo); err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		return
	}

	result.Status = PatchToDo.Status
	result.Details = PatchToDo.Details
	handler.ToDoRepository.Update(&result)
	c.JSON(http.StatusOK, result)
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	var result entities.ToDo
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = handler.ToDoRepository.FindByID(&result, int(id))
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusNotFound, nil)
		return
	}
	handler.ToDoRepository.Delete(&result)
	c.JSON(http.StatusNoContent, gin.H{"message": "To Do Deleted"})
}

func (handler *ToDoHandler) GetToDos(c *gin.Context) {
	var result = []*entities.ToDo{nil}
	ToDos, err := handler.ToDoRepository.FindAll(result)
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if len(ToDos) == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(200, ToDos)
}

func (handler *ToDoHandler) GetToDo(c *gin.Context) {
	var result entities.ToDo
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, err = handler.ToDoRepository.FindByID(&result, int(id))
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println("")
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(200, result)
}
