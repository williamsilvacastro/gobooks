package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Substitua as credenciais pelo seu setup
	var user = "root"
	var password = "admin"
	var host = "localhost"
	var port = "3306"
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/biblioteca", user, password, host, port))
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Erro ao pingar o banco de dados:", err)
	}

	fmt.Println("Conexão ao banco de dados bem-sucedida")
}
