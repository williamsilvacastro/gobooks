package controllers

import (
	"bytes"
	"crud-livros/config"
	"crud-livros/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r)
	return r
}

func TestCreateLivro(t *testing.T) {
	config.Connect()
	router := setupRouter()

	livro := models.Livro{Titulo: "Novo Livro", Autor: "Autor Desconhecido"}
	jsonValue, _ := json.Marshal(livro)
	req, _ := http.NewRequest("POST", "/livros", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateLivro(t *testing.T) {
	config.Connect()
	router := setupRouter()

	livro := models.Livro{Titulo: "Livro Atualizado", Autor: "Autor Atualizado"}
	jsonValue, _ := json.Marshal(livro)
	req, _ := http.NewRequest("PUT", "/livros/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteLivro(t *testing.T) {
	config.Connect()
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/livros/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
