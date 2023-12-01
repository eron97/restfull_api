// config/database/database.go

package database

import (
	"database/sql"
	"log"

	"github.com/eron97/restfull_api.git/config/models"
)

type TodoRepository interface {
	GetAllTasksRepository() ([]models.TodoItem, error)
}

type TodoDB struct {
	db *sql.DB
}

func NewTodoDB(credentials string) (*TodoDB, error) {
	db, err := sql.Open("mysql", credentials)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("[InitDB] Conexão bem-sucedida com AWS RDS")
	return &TodoDB{db}, nil
}

func (tdb *TodoDB) CloseDB() {
	if tdb.db != nil {
		tdb.db.Close()
	}
}

func (tdb *TodoDB) GetAllTasksRepository() ([]models.TodoItem, error) {
	rows, err := tdb.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var todoItems []models.TodoItem

	for rows.Next() {
		var todoitem models.TodoItem
		err := rows.Scan(&todoitem.ID, &todoitem.Task_Name, &todoitem.Priority)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		todoItems = append(todoItems, todoitem)
	}

	return todoItems, nil
}
