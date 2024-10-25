// package main

// import (
//     "crud-livros/config"
//     "crud-livros/controllers"
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     // Conectar ao banco de dados
//     config.Connect()

//     // Inicializar o router do Gin
//     router := gin.Default()

//     // Definir rotas
//     router.GET("/livros", controllers.GetLivros)
//     router.POST("/livros", controllers.CreateLivro)
//     router.PUT("/livros/:id", controllers.UpdateLivro)
//     router.DELETE("/livros/:id", controllers.DeleteLivro)

//     // Iniciar o servidor na porta 8080
//     router.Run(":8080")
// }
