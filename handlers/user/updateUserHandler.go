package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func UpdateUserHandler(c *gin.Context) {
	req := UpdateUserRequest{}
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" && req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one field is required"})
		return
	}

	if req.Email != "" {
		if err := validate.Var(req.Email, "email"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		}
	}

	if req.Name != "" {
		if err := validate.Var(req.Email, "min=4,max=100"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		}
	}

	user := models.User{}

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"data":    user,
	})
}
