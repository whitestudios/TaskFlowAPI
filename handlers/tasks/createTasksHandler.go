package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func CreateTasksHandler(c *gin.Context) {
	var req CreateTasksRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task JSON"})
		return
	}

	if err := validate.Struct(req); err != nil {
		errors := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	newTask := models.Task{
		Title:     req.Title,
		SubTitle:  req.Subtitle,
		Completed: false,
		UserID:    req.UserID,
	}

	if err := db.Create(&newTask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    newTask,
	})
}
