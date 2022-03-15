package handler

import (
	"go-todo-api/entities"
	"go-todo-api/models"
	"go-todo-api/repository"
	"go-todo-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(UserRepo repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{UserRepository: UserRepo}
}

func (handler *UserHandler) PostUser(c *gin.Context) {
	var requestUser = &models.UserRequest{}

	if err := c.BindJSON(&requestUser); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !requestUser.Validate() {
		zap.S().Error("Error: Paswords do not match")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	hashed := utils.HashPassword(requestUser.Password)

	var newUser = &entities.User{
		Email:    requestUser.Email,
		Password: hashed,
		Status: "Pending", //test için
	}

	err := handler.UserRepository.Insert(newUser)
	if err != nil {
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var userResponse = &models.UserResponse{
		Id:        newUser.Id,
		Email:     newUser.Email,
		Status:    newUser.Status,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	}

	c.JSON(http.StatusCreated, userResponse)
}

func (handler *UserHandler) PatchUser(c *gin.Context) {
	var user entities.User
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	user.Id = id

	record, err := handler.UserRepository.FindByID(&user)
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

	var requestUser = &models.UserPatchRequest{}
	if err := c.BindJSON(&requestUser); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !requestUser.Validate() {
		zap.S().Error("Error: Paswords do not match")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	hash := utils.HashPassword(requestUser.Password)
	user.Password = hash

	err = handler.UserRepository.Update(&user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var userResponse = &models.UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var deleteUser = &entities.User{
		Id:     userId,
		Status: "Passive",
	}

	err = handler.UserRepository.Update(deleteUser)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = handler.UserRepository.Delete(deleteUser)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	var user entities.User
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	user.Id = id
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	record, err := handler.UserRepository.FindByID(&user)
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

	var userResponse = &models.UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	var user []*entities.User
	toDos, err := handler.UserRepository.FindAll(user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var toDosResponse = make([]models.UserResponse, len(toDos))

	for i := 0; i < len(toDos); i++ {
		toDosResponse[i] = models.UserResponse{
			Id:        toDos[i].Id,
			Email:     toDos[i].Email,
			Status:    toDos[i].Status,
			CreatedAt: toDos[i].CreatedAt,
			UpdatedAt: toDos[i].UpdatedAt,
			DeletedAt: toDos[i].DeletedAt,
		}
	}
	c.JSON(http.StatusOK, toDosResponse)
}