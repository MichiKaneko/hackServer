package controllers

import (
	"net/http"

	"github.com/MichiKaneko/hackServer/database"
	"github.com/MichiKaneko/hackServer/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var user models.User
	email := c.MustGet("email").(string)
	user_record := database.DB.Where("email = ?", email).First(&user)

	if user_record.Error != nil {
		c.JSON(400, gin.H{"error": user_record.Error.Error()})
		c.Abort()
		return
	}

	var post models.Post
	post.UserID = user.ID

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.DB.Create(&post)
	if record.Error != nil {
		c.JSON(400, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON((http.StatusCreated), gin.H{"id": post.ID, "title": post.Title, "content": post.Content, "user_id": post.UserID})
}

func GetPostByID(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	record := database.DB.First(&post, id)
	if record.Error != nil {
		c.JSON(400, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": post.ID, "title": post.Title, "content": post.Content, "user_id": post.UserID})
}

func GetUserPosts(c *gin.Context) {
	var posts []models.Post
	id := c.Param("id")
	record := database.DB.Where("user_id = ?", id).Find(&posts)
	if record.Error != nil {
		c.JSON(400, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetCurrentUserPosts(c *gin.Context) {
	var user models.User
	email := c.MustGet("email").(string)
	user_record := database.DB.Where("email = ?", email).First(&user)

	if user_record.Error != nil {
		c.JSON(400, gin.H{"error": user_record.Error.Error()})
		c.Abort()
		return
	}

	var posts []models.Post
	record := database.DB.Where("user_id = ?", user.ID).Find(&posts)
	if record.Error != nil {
		c.JSON(400, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
