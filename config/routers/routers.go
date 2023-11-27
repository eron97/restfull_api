package routers

import (
	"github.com/eron97/restfull_api.git/config/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetAllTasksController)
	router.GET("/tasks/:task_name", controllers.GetTaskByName)
	router.POST("/tasks", controllers.AddTask)
}
