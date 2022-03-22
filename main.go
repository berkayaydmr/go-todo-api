package main

import (
	"github.com/go-redis/redis"
	"go-todo-api/common"
	"go-todo-api/handler"
	"go-todo-api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	env := common.GetEnvironment()

	logger := common.NewLogger(env.Debug)
	logger.Info("logger initialized")

	redisClient := redis.NewClient(&redis.Options{
		Addr: env.RedisUrl,
	})

	db := common.ConnectDB(env.DatabaseUrl)

	ToDoRepository := repository.NewToDoRepository(db)
	UserRepository := repository.NewUserRepository(db)
	RedisRepository := repository.NewRedisRepository(redisClient)
	ToDoHandler := handler.NewToDoHandler(ToDoRepository, RedisRepository)
	UserHandler := handler.NewUserHandler(UserRepository, RedisRepository)

	router := gin.Default()
	router.Use(gin.Recovery())

	router.POST("users/login", UserHandler.LoginUser)
	router.POST("users/logout", UserHandler.LogOut)
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
