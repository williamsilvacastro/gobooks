package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Substitua as credenciais pelo seu setup
	var user = getEnv("DB_USER", "root")
	var password = getEnv("DB_PASSWORD", "admin")
	var host = getEnv("DB_HOST", "localhost")
	var port = getEnv("DB_PORT", "3306")

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/biblioteca", user, password, host, port))
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Erro ao pingar o banco de dados:", err)
	}
	// Verifica se a tabela 'livros' existe e cria se não existir
	createTableIfNotExists()
	fmt.Println("Conexão ao banco de dados bem-sucedida")
}

// getEnv retorna o valor da variável de ambiente ou um valor padrão se não estiver definida
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// createTableIfNotExists verifica se a tabela 'livros' existe e a cria se não existir
func createTableIfNotExists() {
	query := `
	CREATE TABLE IF NOT EXISTS livros (
		id INT AUTO_INCREMENT PRIMARY KEY,
		titulo VARCHAR(255) NOT NULL,
		autor VARCHAR(255) NOT NULL,
		descricao TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Erro ao criar a tabela 'livros':", err)
	}
}
