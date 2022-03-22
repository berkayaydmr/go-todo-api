package handler

import (
	"go-todo-api/entities"
	"go-todo-api/models"
	"go-todo-api/repository"
	"go-todo-api/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserRepository  repository.UserRepositoryInterface
	RedisRepository repository.RedisClientInterface
}

func NewUserHandler(UserRepo repository.UserRepositoryInterface, redisRepository repository.RedisClientInterface) *UserHandler {
	return &UserHandler{UserRepository: UserRepo, RedisRepository: redisRepository}
}

func (handler *UserHandler) LoginUser(c *gin.Context) {
	var userRequest = &models.UserLogRequest{}
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var user = entities.User{
		Email: userRequest.Email,
	}

	userResponse, err := handler.UserRepository.FindByEmail(&user)

	if utils.HashPassword(userRequest.Password) != userResponse.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong email or password",
		})
		return
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	redisKey := strconv.FormatUint(userResponse.Id, 10)
	err = handler.RedisRepository.SetData(redisKey, token, time.Minute)
	if err != nil {
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
}

func (handler UserHandler) LogOut(c *gin.Context) {
	var userRequest = &models.UserLogRequest{}
	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var user = entities.User{
		Email: userRequest.Email,
	}

	_, err := handler.UserRepository.FindByEmail(&user)

	userID := strconv.FormatUint(user.Id, 10)

	err = handler.RedisRepository.DeleteData(userID)
	if err != nil {
		zap.S().Error("Something went wrong:", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully logged out",
	})
}

func (handler *UserHandler) PostUser(c *gin.Context) {
	var requestUser = &models.UserRequest{}

	if err := c.BindJSON(&requestUser); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !requestUser.Validate() {
		zap.S().Error("Error: Passwords do not match")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	hashed := utils.HashPassword(requestUser.Password)

	var newUser = &entities.User{
		Email:    requestUser.Email,
		Password: hashed,
		Status:   "Pending", //test için
	}

	err := handler.UserRepository.Insert(newUser)
	if err != nil {
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, utils.UserApiResponse(newUser))
}

func (handler *UserHandler) PatchUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	user := &entities.User{Id: id}

	user, err = handler.UserRepository.FindByID(user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if user == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var requestUser = &models.UserPatchRequest{}
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !requestUser.Validate() {
		zap.S().Error("Error: Passwords do not match or Null variable sent")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if utils.HashPassword(requestUser.OldPassword) != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Current Password not matched"})
		return
	}

	hash := utils.HashPassword(requestUser.Password)
	user.Password = hash

	err = handler.UserRepository.Update(user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, utils.UserApiResponse(user))
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user := entities.User{Id: userId}
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

	user.Status = "PASSIVE"

	err = handler.UserRepository.Update(&user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	user := entities.User{Id: id}
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userResponse, err := handler.UserRepository.FindByID(&user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if userResponse == nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, utils.UserApiResponse(&user))
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
