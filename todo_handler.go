package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addGroup(r *gin.Engine, prefix string) {
	Log.Printf("Adding group for todo handlers")
	v1 := r.Group(prefix + "/todos")
	v1.POST("", createTodo)
	v1.GET("", fetchAllTodo)
	v1.GET("/:id", fetchSingleTodo)
	v1.PUT("/:id", updateTodo)
	v1.DELETE("/:id", deleteTodo)
}

// createTodo add a new todo
func createTodo(c *gin.Context) {
	Log.Printf("Start create todo")
	var tm todoModel
	err := c.BindJSON(&tm)
	if err != nil {
		Log.Printf("Error decoding body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Todo could not be unmarshalled!"})
		return
	}
	Log.Printf("Create todo with title >%s< and complete >%d<", tm.Title, tm.Completed)
	db.Save(&tm)
	c.JSON(http.StatusCreated, tm)
}

// fetchAllTodo fetch all todos
func fetchAllTodo(c *gin.Context) {
	Log.Printf("Start fetch all todos")
	var todos []todoModel
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// fetchSingleTodo fetch a single todo
func fetchSingleTodo(c *gin.Context) {
	Log.Printf("Start fetch single todos")
	var todo todoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Todo with ID not found!"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// updateTodo update a todo
func updateTodo(c *gin.Context) {
	Log.Printf("Start update todo")
	var todo todoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo with ID found!"})
		return
	}
	var tm todoModel
	err := c.BindJSON(&tm)
	if err != nil {
		Log.Printf("Error decoding body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Todo could not be unmarshalled!"})
		return
	}
	todo.Title = tm.Title
	todo.Completed = tm.Completed
	//db.Update(&todo)
	db.Model(&todo).Update("title", tm.Title)
	db.Model(&todo).Update("completed", tm.Completed)
	c.JSON(http.StatusOK, todo)
}

// deleteTodo remove a todo
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo with ID found!"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
