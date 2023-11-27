package database

import (
	"database/sql"
	"log"

	"github.com/eron97/restfull_api.git/config/models"
)

type DBConnector interface {
	Connect() (*sql.DB, error)
}

type MySQLConnector struct {
	Credentials string
}

func (m *MySQLConnector) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.Credentials)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("[InitDB] Conexão bem-sucedida com AWS RDS")
	return db, nil
}

func GetAllTasksRepository(db *sql.DB) ([]models.TodoItem, error) {
	rows, err := db.Query("SELECT * FROM tasks")
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
