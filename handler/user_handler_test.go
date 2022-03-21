package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-todo-api/entities"
	"go-todo-api/mocks"
	"go-todo-api/models"
	"go-todo-api/utils"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoginUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedisRepository := mocks.RedisClientInterface{}

	var userEmail = &entities.User{
		Email: "email",
	}

	var user = &entities.User{
		Id:       0,
		Email:    "email",
		Password: utils.HashPassword(""),
		Status:   "Pending",
	}

	userID := strconv.FormatUint(user.Id, 10)
	token, _ := utils.GenerateToken(user.Id)

	mockUserRepository.On("FindByEmail", userEmail).Return(user, nil)
	mockRedisRepository.On("SetData", userID, token, time.Minute).Return(nil)

	loginRequest := &models.UserLogRequest{Email: "email", Password: ""}
	bin, _ := json.Marshal(loginRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/users/login", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository, &mockRedisRepository).LoginUser(c)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestLogOutUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedisRepository := mocks.RedisClientInterface{}

	var userRequest = &models.UserLogRequest{
		Email:    "email",
		Password: "",
	}

	var userFind = &entities.User{
		Email: "email",
	}

	userPassword := ""
	var user = &entities.User{
		Id:       0,
		Email:    "email",
		Password: utils.HashPassword(userPassword),
	}

	userID := strconv.FormatUint(user.Id, 10)

	mockUserRepository.On("FindByEmail", userFind).Return(nil, nil)
	mockRedisRepository.On("DeleteData", userID).Return(nil)

	bin, _ := json.Marshal(userRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/users/logout", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository, &mockRedisRepository).LogOut(c)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestPostUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := mocks.User()
	userRequest := &models.UserRequest{Email: "email", Password: "", PasswordConfirm: ""}
	user.Password = utils.HashPassword(user.Password)
	userResponse := mocks.UserModelResponse()

	mockUserRepository.On("Insert", user).Return(nil)

	bin, _ := json.Marshal(userRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository, &mockRedis).PostUser(c)

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
	mockRedis := mocks.RedisClientInterface{}

	user := mocks.User()
	userRequest := &models.UserRequest{Email: "email", Password: " ", PasswordConfirm: ""}
	utils.HashPassword(user.Password)

	mockUserRepository.On("Insert", user).Return(nil)

	bin, _ := json.Marshal(userRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository, &mockRedis).PostUser(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestPatchUser_OK(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	userPatchRequest := models.UserPatchRequest{Password: "1111", PasswordConfirm: "1111", OldPassword: "1010"}
	userUpdate := &entities.User{Id: 0, Password: utils.HashPassword("1111")}
	user := &entities.User{Id: 0}
	currentPassword := utils.HashPassword(userPatchRequest.OldPassword)
	userFind := &entities.User{Id: 0, Password: currentPassword}

	mockUserRepository.On("FindByID", user).Return(userFind, nil)
	mockUserRepository.On("Update", userUpdate).Return(nil)

	bin, _ := json.Marshal(userPatchRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("PUT", "/users/:user_id", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository, &mockRedis).PatchUser(c)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body *models.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}

	assert.Equal(t, utils.UserApiResponse(user), body)

}

func TestPatchUser_Fail(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}
	userPatchRequest := &models.UserPatchRequest{Password: "1111", PasswordConfirm: "1111"}
	user := &entities.User{Id: 0}

	mockUserRepository.On("FindByID", user).Return(user, nil)
	mockUserRepository.On("Update", user).Return(nil)

	bin, _ := json.Marshal(userPatchRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("PUT", "/users/:user_id", bytes.NewBuffer(bin))
	NewUserHandler(&mockUserRepository, &mockRedis).PatchUser(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestDeleteUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := &entities.User{Id: 0}
	userUpdated := &entities.User{Id: 0, Status: "PASSIVE"}

	mockUserRepository.On("FindByID", user).Return(user, nil)
	mockUserRepository.On("Update", userUpdated).Return(nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("DELETE", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository, &mockRedis).DeleteUser(c)

	assert.Equal(t, http.StatusNoContent, recorder.Result().StatusCode)
}

func TestDeleteUser_Fail(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := &entities.User{Id: 0}
	userUpdated := &entities.User{Id: 0, Status: "PASSIVE"}

	mockUserRepository.On("FindByID", user).Return(user, nil)
	mockUserRepository.On("Update", userUpdated).Return(nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "a"}}
	c.Request = httptest.NewRequest("DELETE", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository, &mockRedis).DeleteUser(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestGetUser(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := &entities.User{Id: 0}
	userResponse := &models.UserResponse{Id: 0}

	mockUserRepository.On("FindByID", user).Return(user, nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository, &mockRedis).GetUser(c)

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
	mockRedis := mocks.RedisClientInterface{}

	user := &entities.User{Id: 0}

	mockUserRepository.On("FindByID", user).Return(nil, nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/users/:user_id", nil)
	NewUserHandler(&mockUserRepository, &mockRedis).GetUser(c)

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
}

func TestGetUsers(t *testing.T) {
	mockUserRepository := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	var users []*entities.User
	usersResponse := []models.UserResponse{}

	mockUserRepository.On("FindAll", users).Return(users, nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/users", nil)
	NewUserHandler(&mockUserRepository, &mockRedis).GetUsers(c)

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
	mockRedis := mocks.RedisClientInterface{}

	var users []*entities.User

	mockUserRepository.On("FindAll", users).Return(nil, errors.New("error"))

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/users", nil)
	NewUserHandler(&mockUserRepository, &mockRedis).GetUsers(c)

	assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
}
