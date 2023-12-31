// package controllers

package controllers

import (
	"net/http"

	"github.com/eron97/restfull_api.git/config/services"
	"github.com/gin-gonic/gin"
)

var todoService services.TodoListService

func SetTodoService(service services.TodoListService) {
	todoService = service
}

func GetAllTasksController(c *gin.Context) {
	todoItems, err := todoService.GetAllTasksService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": todoItems})
}

// GetTaskByName retorna os detalhes de uma tarefa pelo nome
func GetTaskByName(c *gin.Context) {
	// Retorna os detalhes da task como resposta
	c.JSON(http.StatusOK, gin.H{"tasks": ""})
}

func AddTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tasks": ""})
}
