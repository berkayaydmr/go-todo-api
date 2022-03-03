package main

import (
	"go-todo-api/common"
	"go-todo-api/handler"
	"go-todo-api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	env := common.GetEnviroment()
	db := common.ConnectDB(env.DatabaseUrl)
	ToDoRepository := repository.NewToDoRepository(db)
	ToDoHandler := handler.NewToDoHandler(ToDoRepository)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("user/todos", ToDoHandler.GetToDos)
	router.GET("user/todos/:id", ToDoHandler.GetToDo)
	router.POST("user/todos", ToDoHandler.PostToDo)
	router.PATCH("user/todos/:id", ToDoHandler.PatchToDo)
	router.DELETE("user/todos/:id", ToDoHandler.DeleteToDo)

	router.Run("localhost:8080")
}
