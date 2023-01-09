package controllers

import (
	"net/http"

	"github.com/MichiKaneko/hackServer/database"
	"github.com/MichiKaneko/hackServer/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.DB.Create(&user)
	if record.Error != nil {
		c.JSON(400, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON((http.StatusCreated), gin.H{"id": user.ID, "name": user.Name, "email": user.Email})
}
