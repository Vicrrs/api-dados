package main

import (
	"github.com/Vicrrs/api-dados/config"
	"github.com/Vicrrs/api-dados/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	config.ConectarBanco()

	// Criar um novo roteador
	router := gin.Default()

	// Registrar as rotas de produtos
	routes.ProdutoRotas(router)

	// Iniciar um servidor
	router.Run(":8080")
}
