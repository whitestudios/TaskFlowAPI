package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/TaskFlowAPI/handlers/tasks"
	"github.com/whitestudios/TaskFlowAPI/handlers/user"
)

func initializeRoutes(router *gin.Engine) {
	BasePath := "/api"

	user.Init()
	tasks.Init()
	api := router.Group(BasePath)
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "Hello World!"})
		})
	}

	// TODO: Implement User routes
	usersGroup := api.Group("/user")
	{
		usersGroup.GET("", user.ShowUserHandler)

		usersGroup.POST("", user.CreateUserHandler)

		usersGroup.PUT("", user.UpdateUserHandler)

		usersGroup.DELETE("", user.DeleteUserHandler)

		usersGroup.GET("/all", user.ListUserHandler)
	}
	// Create task's routes
	task := api.Group("/task")
	{
		task.GET("", tasks.ShowTaskHandler)

		task.POST("", tasks.CreateTasksHandler)

		task.PUT("", tasks.UpdateTaskHandler)

		task.PATCH("/:id/status", tasks.StatusTaskHandler)

		task.DELETE("", tasks.DeleteTaskHandler)

		task.GET("/all", tasks.ListTaskHandler)
	}

}
