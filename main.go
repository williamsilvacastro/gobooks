package main

import (
	"context"
	"crud-livros/config"
	"crud-livros/controllers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginAdapter "github.com/awslabs/aws-lambda-go-api-proxy/gin" // Importando o adaptador
	"github.com/gin-gonic/gin"
)

var ginLambda *ginAdapter.GinLambda

func init() {
	// Configurações de conexão com o banco
	config.Connect()

	// Configurações do Gin
	r := gin.Default()
	controllers.RegisterRoutes(r) // Registrar as rotas

	// Inicializa o adaptador Gin para Lambda
	ginLambda = ginAdapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Inicia o Lambda em vez do servidor HTTP tradicional
	lambda.Start(Handler)
}
