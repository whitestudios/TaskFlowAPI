package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func ShowUserHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing user ID",
		})
		return
	}

	user := models.User{}

	if err := db.First(&user, id).Error; err != nil {
		logger.Errf("Error searching user id in ShowUserHandler: %v", err)
		c.JSON(http.StatusFound, gin.H{"error": fmt.Sprintf("user with id %v not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
