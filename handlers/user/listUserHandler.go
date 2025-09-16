package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func ListUserHandler(c *gin.Context) {
	users := []models.User{}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Preload("Tasks").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error finding all tasks of users: %v", err.Error()),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"users":   users,
	})
}
