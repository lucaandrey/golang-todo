package controllers

import (
	"gin-todoapp/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todo []models.Todo

	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo

	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	var err error

	c.BindJSON(&todo)
	err = models.UpdateATodo(&todo, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, todo)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := models.DeleteATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	} else {
		c.JSON(http.StatusOK, gin.H{"id: " + id: "deleted"})
	}
}
