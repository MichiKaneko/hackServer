package controllers

import (
	"github.com/MichiKaneko/hackServer/auth"
	"github.com/MichiKaneko/hackServer/database"
	"github.com/MichiKaneko/hackServer/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.DB.Where("email = ?", request.Email).First(&user)
	
	if record.Error != nil {
		c.JSON(400, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(400, gin.H{"error": credentialError.Error()})
		c.Abort()
		return
	}


	token, err := auth.GenerateJWT(user.Email, user.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}