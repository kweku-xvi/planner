package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/todo/api/v1/controllers"
	"github.com/kweku-xvi/todo/api/v1/middleware"
	"github.com/kweku-xvi/todo/internal/database"
)

func init() {
	database.InitDB()
}

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/tasks", controllers.CreateTask)
		v1.GET("/tasks", controllers.GetAllTasks)
		v1.GET("/tasks/:id", controllers.GetTaskByID)
		v1.PUT("/tasks/:id", controllers.UpdateTask)
		v1.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.SignIn)
		auth.GET("/user/profile", middleware.CheckAuth, controllers.GetUserProfile)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
