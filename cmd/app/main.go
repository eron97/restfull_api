package main

import (
	"log"
	"os"

	"github.com/eron97/restfull_api.git/config/database"
	"github.com/eron97/restfull_api.git/config/routers"
	"github.com/eron97/restfull_api.git/config/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Carregando variáveis de ambiente a partir do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente")
	}

	// Obtendo as credenciais do banco de dados das variáveis de ambiente
	dbCredentials := os.Getenv("DB_CREDENTIALS")

	dbConnector := &database.MySQLConnector{
		Credentials: dbCredentials,
	}

	db, err := dbConnector.Connect()
	if err != nil {
		return
	}

	_ = &services.MyTodoListService{
		DBConnector: db,
	}

	r := gin.Default()

	routers.SetupTaskRoutes(r)

	r.Run()

	db.Close()

}
