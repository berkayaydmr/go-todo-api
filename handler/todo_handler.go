package handler

import (
	"go-todo-api/entities"
	"go-todo-api/models"
	"go-todo-api/repository"
	"go-todo-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ToDoHandler struct {
	ToDoRepository  repository.ToDoRepositoryInterface
	RedisRepository repository.RedisClientInterface
}

func NewToDoHandler(ToDoRepo repository.ToDoRepositoryInterface, redisRepository repository.RedisClientInterface) *ToDoHandler {
	return &ToDoHandler{ToDoRepository: ToDoRepo, RedisRepository: redisRepository}
}

func (handler *ToDoHandler) PostToDo(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	redisToken := handler.RedisRepository.GetData(c.Param("user_id"))
	splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if redisToken != splitToken[1] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tokens not match or expired Token",
		})
		return
	}

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

	c.JSON(http.StatusCreated, utils.ToDoApiResponse(newToDo))
}

func (handler *ToDoHandler) PatchToDo(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	redisToken := handler.RedisRepository.GetData(c.Param("user_id"))
	splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if redisToken != splitToken[1] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tokens not match or expired Token",
		})
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
		UserId: userID,
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

	c.JSON(http.StatusOK, utils.ToDoApiResponse(todo))
}

func (handler *ToDoHandler) DeleteToDo(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	redisToken := handler.RedisRepository.GetData(c.Param("user_id"))
	splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if redisToken != splitToken[1] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tokens not match or expired Token",
		})
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
		UserId: userID,
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
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	redisToken := handler.RedisRepository.GetData(c.Param("user_id"))
	splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if redisToken != splitToken[1] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tokens not match or expired Token",
		})
		return
	}

	var user = &entities.ToDo{
		UserId: userID,
	}

	toDos, err := handler.ToDoRepository.FindAll(user)
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
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	redisToken := handler.RedisRepository.GetData(c.Param("user_id"))
	splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	if redisToken != splitToken[1] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Tokens not match or expired Token",
		})
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
		UserId: userID,
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

	c.JSON(200, utils.ToDoApiResponse(toDo))
}
