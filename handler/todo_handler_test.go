package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-todo-api/entities"
	"go-todo-api/mocks"
	"go-todo-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPostToDo_OK(t *testing.T) {
	mockToDoRepo := mocks.ToDoRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	toDoResponse := mocks.ToDoResponse()
	toDo := mocks.ToDo()

	mockRedis.On("GetData", "0").Return("")
	mockToDoRepo.On("Insert", toDo).Return(nil)

	toDoRequest := models.ToDoRequest{Details: "testDetails", Status: "On Progress"}

	bin, _ := json.Marshal(toDoRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("POST", "/users/:user_id/todos", bytes.NewBuffer(bin))
	c.Request.Header.Add("Authorization", "Bearer ")

	handler := NewToDoHandler(&mockToDoRepo, &mockRedis)
	handler.PostToDo(c)
	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	var body *models.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, toDoResponse, body)
}

func TestPostToDo_FAIL(t *testing.T) {
	mockToDoRepo := mocks.ToDoRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	toDo := mocks.ToDoResponse()

	mockToDoRepo.On("Insert", toDo).Return(nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("POST", "/users/:user_id/todos", nil)
	c.Request.Header.Add("Authorization", "Bearer ")
	handler := NewToDoHandler(&mockToDoRepo, &mockRedis)
	handler.PostToDo(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestPatchToDo_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	toDo := &entities.ToDo{Details: "Updated Detail", Status: "Done"}
	toDoo := &entities.ToDo{}
	toDoPatch := mocks.ToDoPatchResponse()

	mockToDoRepository.On("FindByID", toDoo).Return(toDo, nil)
	mockToDoRepository.On("Update", toDo).Return(nil)
	mockRedis.On("GetData", "0").Return("")

	bin, _ := json.Marshal(toDoPatch)
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}, gin.Param{Key: "todo_id", Value: "0"}}
	c.Request = httptest.NewRequest("PUT", "/users/:user_id/todos/:todo_id", bytes.NewBuffer(bin))
	c.Request.Header.Add("Authorization", "Bearer ")
	handler.PatchToDo(c)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body *models.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}
	assert.Equal(t, toDoPatch, body)
}

func TestPatchToDo_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockUserRepo := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := mocks.User()
	toDo := mocks.ToDoResponse()
	toDoo := &entities.ToDo{}

	mockUserRepo.On("Insert", user).Return(user, nil)
	mockToDoRepository.On("FindByID", toDoo).Return(toDo, nil)
	mockToDoRepository.On("Update", toDo).Return(nil)
	mockRedis.On("GetData", "0").Return(nil)

	var details = "Updated Detail"
	var status = "Done"
	toDoRequest := models.ToDoPatchRequest{Details: &details, Status: &status}

	bin, _ := json.Marshal(toDoRequest)
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "s"}}
	c.Request = httptest.NewRequest("PUT", "/user/todos/0", bytes.NewBuffer(bin))
	c.Request.Header.Add("Authorization", "Bearer ")
	handler.PatchToDo(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestDeleteToDo_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockUserRepo := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := mocks.User()
	toDoo := &entities.ToDo{}

	mockUserRepo.On("Insert", user).Return(user, nil)
	mockToDoRepository.On("FindByID", toDoo).Return(toDoo, nil)
	mockToDoRepository.On("Delete", toDoo).Return(nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}, gin.Param{Key: "todo_id", Value: "0"}}
	c.Request = httptest.NewRequest("DELETE", "/user/todos", nil)
	c.Request.Header.Add("Authorization", "Bearer ")
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	handler.DeleteToDo(c)
	assert.Equal(t, http.StatusNoContent, recorder.Result().StatusCode)
}

func TestDeleteToDo_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockUserRepo := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := mocks.User()
	toDoo := &entities.ToDo{}

	mockUserRepo.On("Insert", user).Return(user, nil)
	mockToDoRepository.On("FindByID", toDoo).Return(toDoo, nil)
	mockToDoRepository.On("Delete", toDoo).Return(nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "a"}}
	c.Request = httptest.NewRequest("DELETE", "/user/todos", nil)
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	handler.DeleteToDo(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestGetToDos_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockUserRepo := mocks.UserRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	user := mocks.User()
	var toDo = &entities.ToDo{
		UserId: user.Id,
	}
	var toDos = []*entities.ToDo{}
	var toDosResponse = []models.ToDo{}

	mockUserRepo.On("Insert", user).Return(user, nil)
	mockToDoRepository.On("FindAll", toDo).Return(toDos, nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/users/:user_id/todos", nil)
	c.Request.Header.Add("Authorization", "Bearer ")
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	handler.GetToDos(c)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body []models.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}
	assert.Equal(t, toDosResponse, body)
}

func TestGetToDos_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	var toDos []*entities.ToDo
	mockToDoRepository.On("FindAll", toDos).Return(toDos, nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/users/:user_id/todos", nil)
	c.Request.Header.Add("Authorization", "Bearer ")
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	handler.GetToDos(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestGetToDo_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	toDo := &entities.ToDo{}
	toDoResponse := models.ToDo{
		CreatedAt: toDo.CreatedAt.String(),
		UpdatedAt: toDo.UpdatedAt.String(),
	}

	mockToDoRepository.On("FindByID", toDo).Return(toDo, nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "user_id", Value: "0"}, gin.Param{Key: "todo_id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/users/:user_id/todos/:todo_id", nil)
	c.Request.Header.Set("Authorization", "Bearer ")
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	handler.GetToDo(c)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body models.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		zap.S().Error("Error: ", zap.Error(err))
	}
	assert.Equal(t, toDoResponse, body)
}

func TestGetToDo_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	mockRedis := mocks.RedisClientInterface{}

	toDo := &entities.ToDo{}
	mockToDoRepository.On("FindByID", toDo).Return(toDo, nil)
	mockRedis.On("GetData", "0").Return("")

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "a"}}
	c.Request = httptest.NewRequest("GET", "/user/todos/0", nil)
	c.Request.Header.Add("Authorization", "Bearer ")
	handler := NewToDoHandler(&mockToDoRepository, &mockRedis)
	handler.GetToDo(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}
