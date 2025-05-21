package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func DeleteUserHandler(c *gin.Context) {
	// Implementation of DeleteUserHandler
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing user ID",
		})
		return
	}

	user := models.User{}

	if err := db.First(&user, id).Error; err != nil {
		logger.Errf("Error searching user id in ShowUserHandler: ", err)
		c.JSON(http.StatusFound, gin.H{"error": fmt.Sprintf("user with id %v not found", id)})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error deleting user with id %v not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "user deleted", "data": user})

}
