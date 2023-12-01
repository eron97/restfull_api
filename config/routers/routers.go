package routers

import (
	"github.com/eron97/restfull_api.git/config/controllers"
	"github.com/eron97/restfull_api.git/config/services"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine, todoService services.TodoListService) {
	controllers.SetTodoService(todoService)

	router.GET("/tasks", controllers.GetAllTasksController)
	router.GET("/tasks/:task_name", controllers.GetTaskByName)
	router.POST("/tasks", controllers.AddTask)
}
