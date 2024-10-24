package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/todo/api/v1/controllers"
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

	r.Run() // listen and serve on 0.0.0.0:8080
}
