package services

import (
	"database/sql"

	"github.com/eron97/restfull_api.git/config/database"
	"github.com/eron97/restfull_api.git/config/models"
)

type TodoListService interface {
	GetAllTasksService() ([]TodoItem, error)
}

type MyTodoListService struct {
	DBConnector *sql.DB
}

type TodoItem struct {
	ID        int    `json:"id"`
	Task_Name string `json:"task_name"`
	Priority  string `json:"priority"`
}

func (s *MyTodoListService) GetAllTasksService() ([]models.TodoItem, error) {
	todoItems, err := database.GetAllTasksRepository(s.DBConnector)
	if err != nil {
		return nil, err
	}

	// Outras etapas de processamento de negócios
	// .
	// .
	// .

	return todoItems, nil
}
