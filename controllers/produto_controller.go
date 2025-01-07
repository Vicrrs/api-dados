package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Vicrrs/api-dados/config"
	"github.com/Vicrrs/api-dados/models"
)

// ObterTodosProdutos retorna todos os produtos
func ObterTodosProdutos(c *gin.Context) {
	var produtos []models.Produto
	config.DB.Find(&produtos)
	c.JSON(http.StatusOK, produtos)
}

// CriarProduto cria um novo produto
func CriarProduto(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
	}
	config.DB.Create(&produto)
	c.JSON(http.StatusCreated, produto)
}

// ObterProdutoPorID retorna um produto pelo ID
func ObterProdutoPorID(c *gin.Context) {
	id := c.Param("id")
	var produto models.Produto
	if err := config.DB.First(&produto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	c.JSON(http.StatusOK, produto)
}

// AtualizarProduto atualiza um produto pelo ID
func AtualizarProduto(c *gin.Context) {
	id := c.Param("id")
	var produto models.Produto

	if err := config.DB.First(&produto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&produto)
	c.JSON(http.StatusOK, produto)
}

// DeletarProduto deleta um produto pelo ID
func DeletarProduto(c *gin.Context) {
	id := c.Param("id")
	var produto models.Produto

	if err := config.DB.First(&produto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	config.DB.Delete(&produto)
	c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}
