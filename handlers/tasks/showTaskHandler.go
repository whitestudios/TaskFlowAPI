package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func ShowTaskHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing task ID"})
		return
	}

	task := models.Task{}

	if err := db.First(&task).Error; err != nil {
		logger.Errf("Error searching user id in ShowTaskHandler: %v", err)
		c.JSON(http.StatusFound, gin.H{"error": fmt.Sprintf("task with id: %v not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User find with success",
		"data":    task,
	})
}
