package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Define a struct for the data model
type Todo struct {
ID int `json:"id"`
Title string `json:"title`
Status string `json:"status`
}

var todos []Todo

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtSecretKey = []byte("your_secret_key_here")

func main() {
// Initialize the Gin router
router := gin.Default()
// Set up routes for the API
router.GET("/todos", getTodos)
router.POST("/todos", createTodo)
router.GET("/todos/:id", getTodo)
router.PUT("/todos/:id", updateTodo)
router.DELETE("/todos/:id", deleteTodo)

// Start the server
router.Run(":8080")
}

// List all todos
func getTodos(c *gin.Context) {
c.JSON(http.StatusOK, todos)
}

// Create a new todo
func createTodo(c *gin.Context) {
var todo Todo
c.BindJSON(&todo)
// Simulate ID generation
todo.ID = len(todos) + 1
todos = append(todos, todo)
c.JSON(http.StatusCreated, todo)
}

// Get a single todo by ID
func getTodo(c *gin.Context) {
id := c.Param("id")
for _, todo := range todos {
	if id == strconv.Itoa(todo.ID) {
		c.JSON(http.StatusOK, todo)
		return
	}
}
c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// Update a todo by ID
func updateTodo(c *gin.Context) {
var updatedTodo Todo
c.BindJSON(&updatedTodo)
id := c.Param("id")

for i, todo := range todos {
	if id == strconv.Itoa(todo.ID) {
		todos[i].Title = updatedTodo.Title
		todos[i].Status = updatedTodo.Status
		c.JSON(http.StatusOK, todos[i])
		return
	}
}
c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// Delete a todo by ID
func deleteTodo(c *gin.Context) {
id := c.Param("id")
for i, todo := range todos {
	if id == strconv.Itoa(todo.ID) {
		todos = append(todos[:i], todos[i+1:]...)
		c.JSON(http.StatusOK, gin.H{"result": "Todo deleted"})
		return
	}
}
c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}