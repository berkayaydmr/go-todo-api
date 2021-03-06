package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-todo-api/entities"
	"go-todo-api/mocks"
	"go-todo-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostToDo_OK(t *testing.T) {
	mockToDoRepo := mocks.ToDoRepositoryInterface{}
	toDo := mocks.ToDoResponse()
	mockToDoRepo.On("Insert", toDo).Return(nil)

	toDoRequest := models.ToDoRequest{Details: "testDetails", Status: "On Progress"}

	bin, _ := json.Marshal(toDoRequest)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/user/todos", bytes.NewBuffer(bin))
	handler := NewToDoHandler(&mockToDoRepo)
	handler.PostToDo(c)
	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	var body *entities.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, toDo, body)
}

func TestPostToDo_FAIL(t *testing.T) {
	t.Run("POST To-do", func(t *testing.T) {
		mockToDoRepo := mocks.ToDoRepositoryInterface{}
		toDo := mocks.ToDoResponse()
		mockToDoRepo.On("Insert", toDo).Return(toDo, nil)

		recorder := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(recorder)
		handler := NewToDoHandler(&mockToDoRepo)
		handler.PostToDo(c)

		assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	})
}

func TestPatchToDo_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	toDo := mocks.ToDoResponse()
	toDoo := &entities.ToDo{}
	toDoPatch := mocks.ToDoPatchResponse()
	mockToDoRepository.On("FindByID", toDoo).Return(toDo, nil)
	mockToDoRepository.On("Update", toDoPatch).Return(toDoPatch, nil)

	var details string = "Updated Detail"
	var status string = "Done"
	toDoRequest := models.ToDoPatchRequest{Details: &details, Status: &status}

	bin, _ := json.Marshal(toDoRequest)
	handler := NewToDoHandler(&mockToDoRepository)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "0"}}
	c.Request = httptest.NewRequest("PUT", "/user/todos/0", bytes.NewBuffer(bin))
	handler.PatchToDo(c)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body *entities.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, toDoPatch, body)

}

func TestPatchToDo_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	toDoo := &entities.ToDo{}
	toDoPatch := models.ToDoRequest{}
	mockToDoRepository.On("FindByID", toDoo).Return(toDoo, nil)
	mockToDoRepository.On("Update", toDoPatch).Return(toDoPatch, nil)

	var details string = "Updated Detail"
	var status string = "Done"
	toDoRequest := models.ToDoPatchRequest{Details: &details, Status: &status}

	bin, _ := json.Marshal(toDoRequest)
	handler := NewToDoHandler(&mockToDoRepository)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "s"}}
	c.Request = httptest.NewRequest("PUT", "/user/todos/0", bytes.NewBuffer(bin))
	handler.PatchToDo(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestDeleteToDo_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	toDoo := &entities.ToDo{}
	mockToDoRepository.On("FindByID", toDoo).Return(toDoo, nil)
	mockToDoRepository.On("Delete", toDoo).Return(nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "0"}}
	c.Request = httptest.NewRequest("DELETE", "/user/todos", nil)
	handler := NewToDoHandler(&mockToDoRepository)
	handler.DeleteToDo(c)
	assert.Equal(t, http.StatusNoContent, recorder.Result().StatusCode)
}

func TestDeleteToDo_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	toDoo := &entities.ToDo{}
	mockToDoRepository.On("FindByID", toDoo).Return(toDoo, nil)
	mockToDoRepository.On("Delete", toDoo).Return(nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "a"}}
	c.Request = httptest.NewRequest("DELETE", "/user/todos", nil)
	handler := NewToDoHandler(&mockToDoRepository)
	handler.DeleteToDo(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestGetToDos_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	var toDos []*entities.ToDo

	mockToDoRepository.On("FindAll", toDos).Return(toDos, nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/user/todos", nil)
	handler := NewToDoHandler(&mockToDoRepository)
	handler.GetToDos(c)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body []*entities.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, toDos, body)
}

func TestGetToDos_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	var toDos []*entities.ToDo
	err := errors.New("err")
	mockToDoRepository.On("FindAll", toDos).Return(nil, err)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/user/todos", nil)
	handler := NewToDoHandler(&mockToDoRepository)
	handler.GetToDos(c)
	assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
}

func TestGetToDo_OK(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	toDo := &entities.ToDo{}
	mockToDoRepository.On("FindByID", toDo).Return(toDo, nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "0"}}
	c.Request = httptest.NewRequest("GET", "/user/todos/0", nil)
	handler := NewToDoHandler(&mockToDoRepository)
	handler.GetToDo(c)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)

	var body *entities.ToDo
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		fmt.Print(err)
	}
	assert.Equal(t, toDo, body)
}

func TestGetToDo_Fail(t *testing.T) {
	mockToDoRepository := mocks.ToDoRepositoryInterface{}
	toDo := &entities.ToDo{}
	mockToDoRepository.On("FindByID", toDo).Return(toDo, nil)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "a"}}
	c.Request = httptest.NewRequest("GET", "/user/todos/0", nil)
	handler := NewToDoHandler(&mockToDoRepository)
	handler.GetToDo(c)
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}
