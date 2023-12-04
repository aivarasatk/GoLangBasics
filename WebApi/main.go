package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{
		Id:        "1",
		Item:      "Clean room",
		Completed: false,
	},
	{
		Id:        "2",
		Item:      "Read",
		Completed: false,
	},
	{
		Id:        "3",
		Item:      "Dance",
		Completed: false,
	},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getSpecificTodo(context *gin.Context) {
	id := context.Param("id")

	todo, errorResponse := getTodoById(id)

	if errorResponse != nil {
		context.JSON(http.StatusNotFound, errorResponse.Error())
	} else {
		context.JSON(http.StatusOK, todo)
	}

}

func getTodoById(id string) (*Todo, error) {
	for _, todo := range todos {
		if todo.Id == id {
			return &todo, nil
		}
	}

	return nil, errors.New(strconv.FormatInt(http.StatusNotFound, 10))
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getSpecificTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:8080")

}
