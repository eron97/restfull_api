// main.go

package main

import (
	"log"
	"os"

	"github.com/eron97/restfull_api.git/config/database"
	"github.com/eron97/restfull_api.git/config/routers"
	"github.com/eron97/restfull_api.git/config/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente:", err)
	}

	dbCredentials := os.Getenv("DB_CREDENTIALS")

	tdb, err := database.NewTodoDB(dbCredentials)
	if err != nil {
		log.Fatal("Erro ao inicializar o banco de dados:", err)
	}

	defer tdb.CloseDB()

	todoService := services.NewMyTodoListService(tdb)
	r := gin.Default()
	routers.SetupTaskRoutes(r, todoService)
	r.Run()
}
