package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iftekhersunny/restful_api_with_tdd/models"
)

////////////////////////////////////////////
// Todos controller
////////////////////////////////////////////
type TodosController struct{}

// Get all todos list
func (tc *TodosController) Index(c *gin.Context) {
	todo := models.NewTodo()

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"todos": todo.Get(),
		},
		"status": http.StatusOK,
	})
}

// Create a new todo
func (tc *TodosController) Create(c *gin.Context) {
	name := c.PostForm("name")

	todo := models.NewTodo()

	if todo.Insert(name) {
		c.JSON(http.StatusCreated, gin.H{
			"message":     "New todo has been added",
			"status":      http.StatusCreated,
			"total_todos": todo.Count(),
		})
	}
}

// Get a todo by the given todo id
func (tc *TodosController) Get(c *gin.Context) {
	id := c.Param("id")

	todo := models.NewTodo()
	todoId, _ := strconv.Atoi(id)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"todo": todo.Find(todoId),
		},
		"status": http.StatusOK,
	})
}

// Update an existing todo
func (tc *TodosController) Update(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")

	todo := models.NewTodo()
	todoId, _ := strconv.Atoi(id)

	if todo.Update(todoId, name) {
		c.JSON(http.StatusOK, gin.H{
			"message": "New todo has been updated",
			"status":  http.StatusOK,
		})
	}
}

// Delete an existing todo by the given todo id
func (tc *TodosController) Delete(c *gin.Context) {
	id := c.Param("id")

	todo := models.NewTodo()
	todoId, _ := strconv.Atoi(id)

	if todo.Delete(todoId) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "New todo has been deleted",
			"status":      http.StatusOK,
			"total_todos": todo.Count(),
		})
	}
}
