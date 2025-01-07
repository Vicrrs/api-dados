package models

import (
	"gorm.io/gorm"
)

// Produto representa um produto na área de dados
type Produto struct {
	gorm.Model
	Nome       string  `json:"nome" binding:"required"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco" binding:"required"`
	Quantidade int     `json:"quantidade" binding:"required"`
}
