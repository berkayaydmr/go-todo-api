package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-todo-api/entities"
	"go-todo-api/mocks"
	"go-todo-api/models"
	"go-todo-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPostUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	user := mocks.User()
	userRequest := &models.UserRequest{Email: "email", Password: "", PasswordConfirm: ""}
	password := utils.HashPassword(user.Password)
	user.Password = password
	userResponse := mocks.UserModelResponse()

	mockUserRepository.On("Insert", user).Return(nil)

	bin, _ := json.Marshal(userRequest)
	fmt.Print(bin)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository).PostUser(c)

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	var body *models.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}

	assert.Equal(t, userResponse, body)
}

func TestPostUser_FAIL(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	user := mocks.User()
	userRequest := &models.UserRequest{Email: "email", Password: " ", PasswordConfirm: ""}
	password := utils.HashPassword(user.Password)
	user.Password = password

	mockUserRepository.On("Insert", user).Return(nil)

	bin, _ := json.Marshal(userRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository).PostUser(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestPatchUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	userPatchRequest := &models.UserPatchRequest{Password: "1111", PasswordConfirm: "1111"}
	userUpdate := &entities.User{Id: 0, Password: utils.HashPassword(userPatchRequest.Password)}
	user := &entities.User{Id: 0}

	mockUserRepository.On("FindByID", user).Return(user, nil)
	mockUserRepository.On("Update", userUpdate).Return(nil)

	bin, _ := json.Marshal(userPatchRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("PUT", "/users/:user_id", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository).PatchUser(c)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestPatchUser_Fail(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	userPatchRequest := &models.UserPatchRequest{Password: "111", PasswordConfirm: "2222"}
	user := &entities.User{Id: 0}

	mockUserRepository.On("FindByID", user).Return(user, nil)
	mockUserRepository.On("Update", user).Return(nil)

	bin, _ := json.Marshal(userPatchRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("PUT", "/users/:user_id", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository).PatchUser(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestDeleteUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	user := &entities.User{Id: 0, Status: "Passive"}

	mockUserRepository.On("Update", user).Return(nil)
	mockUserRepository.On("Delete", user).Return(nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("DELETE", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository).DeleteUser(c)

	assert.Equal(t, http.StatusNoContent, recorder.Result().StatusCode)
}

func TestDeleteUser_Fail(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	user := &entities.User{Id: 0, Status: "Passive"}

	mockUserRepository.On("Update", user).Return(nil)
	mockUserRepository.On("Delete", user).Return(nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "a"}}
	c.Request = httptest.NewRequest("DELETE", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository).DeleteUser(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestGetUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	user := &entities.User{Id: 0}
	userResponse := &models.UserResponse{Id: 0}

	mockUserRepository.On("FindByID", user).Return(user,nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository).GetUser(c)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body *models.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}

	assert.Equal(t, userResponse, body)
}

func TestGetUser_Fail(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	user := &entities.User{Id: 0}

	mockUserRepository.On("FindByID", user).Return(nil,nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository).GetUser(c)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestGetUsers(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	var users []*entities.User
	usersResponse := []models.UserResponse{}

	mockUserRepository.On("FindAll", users).Return(users,nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/users", nil)
	NewUserHandler(&mockUserRepository).GetUsers(c)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body []models.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}

	assert.Equal(t, usersResponse, body)
}

func TestGetUsers_Fail(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	var users []*entities.User

	mockUserRepository.On("FindAll", users).Return(nil,errors.New("error"))

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/users", nil)
	NewUserHandler(&mockUserRepository).GetUsers(c)

	assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
}