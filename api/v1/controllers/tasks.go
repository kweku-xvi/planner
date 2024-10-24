package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/todo/api/v1/models"
	"github.com/kweku-xvi/todo/internal/database"
)

var body struct {
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Priority    string    `json:"priority" form:"priority"`
	Deadline    time.Time `json:"deadline" form:"deadline"`
	Status      string    `json:"status" form:"status"`
}

func CreateTask(c *gin.Context) {
	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error":err.Error(),
		})
		return
	}

	task := models.Task{
		Title: body.Title,
		Description: body.Description,
		Priority: body.Priority,
		Deadline: body.Deadline,
		Status: body.Status,
	}

	result := database.DB.Create(&task)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error":result.Error,
		})
		return
	}

	c.JSON(201, gin.H{
		"success":true,
		"task":task,
	})
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	
	result := database.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"error":"task not found",
		})
		return
	} 

	c.JSON(200, gin.H{
		"success":true,
		"task":task,
	})
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	results := database.DB.Find(&tasks)
	if results.Error != nil {
		c.JSON(400, gin.H{
			"error":results.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success":true,
		"tasks":tasks,
	})

}

func UpdateTask(c *gin.Context) {
	var task models.Task

	id := c.Param("id")

	result := database.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"error":"task not found",
		})
		return
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error":err.Error(),
		})
		return
	}

	database.DB.Model(&task).Updates(models.Task{
		Title: body.Title,
		Description: body.Description,
		Status: body.Status,
		Priority: body.Priority,
		Deadline: body.Deadline,
	})
	c.JSON(200, gin.H{
		"success":true,
		"message":"task updated successfully",
		"task":task,
	})
}

func DeleteTask(c *gin.Context) {
	var task models.Task

	id := c.Param("id")

	result := database.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message":"task not found",
		})
		return
	}

	database.DB.Delete(&task, id)
	c.Status(204)
}