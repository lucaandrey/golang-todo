package controllers

import (
	"gin-todoapp/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var u []models.User
	c.BindJSON(&u)
	err := models.GetAllUsers(&u)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, u)
	}

}

func CreateAnUser(c *gin.Context) {
	var u models.User
	c.BindJSON(u)
	err := models.CreateUser(&u)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusCreated, u)
	}
}

func GetAUser(c *gin.Context) {
	var u models.User
	id := c.Params.ByName("id")
	err := models.GetAnUser(id, &u)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func UpdateAUser(c *gin.Context) {
	var u models.User
	id := c.Params.ByName("id")
	c.BindJSON(&u)
	err := models.UpdateUser(id, &u)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, u)
	}

}

func DeleteAUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var u models.User
	err := models.DeleteUser(id, &u)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"Id: " + id: "deleted"})
	}
}
