package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	BasePath := "/api"

	api := router.Group(BasePath)
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "Hello World!"})
		})
	}
}
