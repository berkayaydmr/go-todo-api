package main

import (
	"go-todo-api/common"
	"go-todo-api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/todos",getTodos)
	router.POST("/todos", createToDo)

	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	var ToDo entities.ToDo
	db := common.ConnectDB()
	c.IndentedJSON(http.StatusOK, db.First(&ToDo, 0))
	
}

func createToDo(c *gin.Context) {
	db := common.ConnectDB()

	toDo := entities.ToDo{Details: "To-do", Status: "On-Progress"}

	db.Create(toDo)
}
