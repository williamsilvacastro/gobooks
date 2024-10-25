package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/livros",
	}

	response, err := Handler(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)
}
