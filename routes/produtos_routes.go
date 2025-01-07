package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Vicrrs/api-dados/controllers"
)

func ProdutoRotas(router *gin.Engine) {
	produtos := router.Group("/produtos")
	{
		produtos.GET("/", controllers.ObterTodosProdutos)
		produtos.POST("/", controllers.CriarProduto)
		produtos.GET("/:id", controllers.ObterProdutoPorID)
		produtos.PUT("/:id", controllers.AtualizarProduto)
		produtos.DELETE("/:id", controllers.DeletarProduto)
	}
}
