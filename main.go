package main

import (
	"github.com/MichiKaneko/hackServer/controllers"
	"github.com/MichiKaneko/hackServer/database"
	"github.com/MichiKaneko/hackServer/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("user:p@ssw0rd@tcp(localhost:3306)/hack?charset=utf8mb4&parseTime=True&loc=Local")
	database.Migrate()

	r := initRouter()
	r.Run(":8080")
}

func initRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/user/:id", controllers.GetUserByID)
		}
		user := api.Group("/users").Use(middlewares.Auth()).Use(middlewares.CurrentUser())
		{
			user.GET("/me", controllers.GetMe)
			user.POST("/post", controllers.CreatePost)
			user.GET("/post", controllers.GetCurrentUserPosts)
		}
	}
	return r
}
