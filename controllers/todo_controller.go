package controllers

import (
	"net/http"

	"todo-api/config"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	result := config.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve todos",
		})

		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	var todo models.Todo
	result := config.DB.First(&todo, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultDB := config.DB.Create(&todo)

	if resultDB.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	if err := config.DB.First(&todo, c.Param("id")); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&todo).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	if err := config.DB.First(&todo, c.Param("id")); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
