package main

import (
	"go-todo-api/common"
	"go-todo-api/handler"
	"go-todo-api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	env := common.GetEnviroment()

	logger := common.NewLogger(env.Debug)
	logger.Info("logger initilazed")

	db := common.ConnectDB(env.DatabaseUrl)
	ToDoRepository := repository.NewToDoRepository(db)
	UserRepository := repository.NewUserRepository(db)
	ToDoHandler := handler.NewToDoHandler(ToDoRepository)
	UserHandler := handler.NewUserHandler(UserRepository)

	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("users", UserHandler.GetUsers)
	router.GET("users/:user_id", UserHandler.GetUser)
	router.POST("users", UserHandler.PostUser)
	router.PATCH("users/:user_id", UserHandler.PatchUser)
	router.DELETE("users/:user_id", UserHandler.DeleteUser)

	router.GET("users/:user_id/todos", ToDoHandler.GetToDos)
	router.GET("users/:user_id/todos/:todo_id", ToDoHandler.GetToDo)
	router.POST("users/:user_id/todos", ToDoHandler.PostToDo)
	router.PATCH("users/:user_id/todos/:todo_id", ToDoHandler.PatchToDo)
	router.DELETE("users/:user_id/todos/:todo_id", ToDoHandler.DeleteToDo)

	router.Run(env.RouterUrl)
}
