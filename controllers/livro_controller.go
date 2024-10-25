package controllers

import (
	"crud-livros/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Buscar todos os livros
func GetLivros(c *gin.Context) {
	livros, err := models.BuscarLivros()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livros"})
		return
	}
	c.JSON(http.StatusOK, livros)
}

// Buscar um livro por ID
func GetLivroByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	livro, err := models.BuscarLivroPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

// Criar um novo livro
func CreateLivro(c *gin.Context) {
	var novoLivro models.Livro
	if err := c.ShouldBindJSON(&novoLivro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := models.CriarLivro(&novoLivro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar livro"})
		return
	}

	c.JSON(http.StatusCreated, novoLivro)
}

// Atualizar um livro existente
func UpdateLivro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var livro models.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	livro.ID = id
	if err := models.AtualizarLivro(&livro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar livro"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

// Deletar um livro
func DeleteLivro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := models.DeletarLivro(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar livro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso"})
}
