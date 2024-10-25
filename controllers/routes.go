package controllers

import "github.com/gin-gonic/gin"

// RegisterRoutes define todas as rotas disponíveis e os controladores associados.
func RegisterRoutes(r *gin.Engine) {
	// Grupo de rotas para o recurso "livros"
	livros := r.Group("/livros")
	{
		livros.GET("/", GetLivros)         // Listar todos os livros
		livros.GET("/:id", GetLivroByID)   // Obter um livro por ID
		livros.POST("/", CreateLivro)      // Criar um novo livro
		livros.PUT("/:id", UpdateLivro)    // Atualizar um livro existente
		livros.DELETE("/:id", DeleteLivro) // Deletar um livro
	}
}
