package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Vicrrs/api-dados/models"
)

// DB é a instancia do banco de dados
var DB *gorm.DB

// ConnectDB conecta ao banco de dados
func ConectarBanco() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar o arquivo .env")
	}

	// pega a string de conexão do banco de dados
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Váriavel de ambiente DATABASE_URL não definida")
	}

	// abre a conexão com o banco de dados
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados")
	}

	// automigrate pra criar as tabelas, se ainda nao existirem
	if err := database.AutoMigrate(&models.Produto{}); err != nil {
		log.Fatal("Falha ao migrar o modelo Produto:", err)
	}

	DB = database
	fmt.Println("Conectado ao banco de dados com sucesso!")
}
