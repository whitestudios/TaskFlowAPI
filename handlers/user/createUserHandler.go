package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/whitestudios/TaskFlowAPI/models"
)

func CreateUserHandler(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json"})
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

	new_user := models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := db.Create(&new_user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    new_user,
	})

}
