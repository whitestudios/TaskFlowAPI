package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func ListTaskHandler(c *gin.Context) {
	var tasks []models.Task

	if err := db.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks find successfully",
		"tasks":   tasks,
	})
}
