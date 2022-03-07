package main

import (
	"go-todo-api/common"
	"go-todo-api/handler"
	"go-todo-api/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("Logger initialize error: ", zap.Error(err))
	}
	zap.ReplaceGlobals(logger)

	err = godotenv.Load("database.env")
	if err != nil {
		zap.S().Error("Error: ", err)
		return
	}
	
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

	router.Run(env.RouterUrl)
}
