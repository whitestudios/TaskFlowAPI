package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func DeleteTaskHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing task ID",
		})
	}

	task := models.Task{}

	if err := db.First(task, id).Error; err != nil {
		logger.Errf("Error searching task id in DeleteTaskHandler: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("task with id %v not found", id)})
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error deleting task: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted",
		"data":    task,
	})
}
