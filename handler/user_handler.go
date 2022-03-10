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

type UserHandler struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserHandler(UserRepo repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{UserRepository: UserRepo}
}

func (handler *UserHandler) PostUser(c *gin.Context){
	var requestUser =  &models.UserRequest{}

	if err := c.ShouldBind(&requestUser); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !requestUser.Validate() {
		zap.S().Error("Error: Paswords do not match")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var newUser = &entities.User{
		Email: requestUser.Email,
		Password: requestUser.Password,
	}

	err := handler.UserRepository.Insert(newUser)
	if err != nil {
		zap.S().Error("Error: ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (handler *UserHandler) PatchUser(c *gin.Context){
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
	if err := c.ShouldBind(&requestUser); err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !requestUser.Validate() {
		zap.S().Error("Error: Paswords do not match")
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if requestUser.Password != nil {
		user.Password = *requestUser.Password
	}

	err = handler.UserRepository.Update(&user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) DeleteUser(c *gin.Context)  {
	userId, err := strconv.ParseUint(c.Param("user_id"),10,64)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, nil)
		return 
	}

	var deleteUser = &entities.User{
		Id: userId,
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

	c.JSON(http.StatusNoContent,nil)
}

func (handler *UserHandler) GetUser(c *gin.Context)  {
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

	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) GetUsers(c *gin.Context)  {
	var user []*entities.User
	toDos, err := handler.UserRepository.FindAll(user)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, toDos)
}