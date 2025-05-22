package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func UpdateTaskHandler(c *gin.Context) {
	var req UpdateTasksRequest

	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task ID missing",
		})
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request json body, this request must be title or subtitle or completed",
		})
	}

	taskToChange := models.Task{}

	if err := db.First(&taskToChange, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("task with id %v not found", id),
		})
	}

	if req.Title != "" {
		taskToChange.Title = req.Title
	}

	if req.Subtitle != "" {
		taskToChange.SubTitle = req.Subtitle
	}

	if req.Completed != nil {
		taskToChange.Completed = *req.Completed
	}

	if err := db.Save(&taskToChange).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error saving task on database: %v", err.Error()),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
		"data":    taskToChange,
	})
}
