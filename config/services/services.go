// services/services.go

package services

import (
	"github.com/eron97/restfull_api.git/config/database"
	"github.com/eron97/restfull_api.git/config/models"
)

type TodoListService interface {
	GetAllTasksService() ([]models.TodoItem, error)
}

type MyTodoListService struct {
	Repository database.TodoRepository
}

func NewMyTodoListService(repository database.TodoRepository) *MyTodoListService {
	return &MyTodoListService{
		Repository: repository,
	}
}

func (ts *MyTodoListService) GetAllTasksService() ([]models.TodoItem, error) {
	return ts.Repository.GetAllTasksRepository()
	// restante do processamento da lógica de negócios
	// output
}
