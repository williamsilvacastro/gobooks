# CRUD de Livros com Gin e MySQL

Este é um projeto de CRUD de livros desenvolvido em Golang, utilizando o framework **Gin** para manipulação de rotas e **MySQL** para armazenamento de dados. O projeto segue a arquitetura MVC (Model-View-Controller) para organização de código.

## Estrutura do Projeto

```bash
crud-livros
├── controllers/
│   └── livro_controller.go       # Controladores das rotas de livros
├── models/
│   └── livro.go                  # Modelo do Livro e funções CRUD
├── config/
│   └── db.go                     # Configuração do banco de dados
└── main.go                       # Configuração do servidor e rotas principais


## Pré-requisitos

- **Golang**: Instale o Go na versão 1.16 ou superior.
- **MySQL**: Instale o MySQL e configure um banco de dados chamado `biblioteca`.

## Instalação

1. Clone este repositório:
   ```bash
   git clone https://github.com/williamsilvacastro/gobooks
   cd gobooks
   ```

2. Inicialize o módulo Go:
   ```bash
   go mod init crud-livros
   ```

3. Instale as dependências:
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/go-sql-driver/mysql
   ```

4. Configure o banco de dados MySQL:
   - Conecte-se ao MySQL e crie o banco de dados e a tabela com o seguinte comando:

     ```sql
     CREATE DATABASE biblioteca;

     USE biblioteca;

     CREATE TABLE livros (
         id INT AUTO_INCREMENT PRIMARY KEY,
         titulo VARCHAR(100),
         autor VARCHAR(100),
         descricao TEXT
     );
     ```

5. Atualize a configuração do banco de dados em `config/db.go` com suas credenciais MySQL.

## Executando o Projeto

Para iniciar o servidor, execute o comando:

```bash
go run main.go
```

O servidor estará disponível na porta **8080**.

## Endpoints

Aqui estão os endpoints disponíveis para o CRUD de livros:

### 1. Listar todos os livros
   - **Rota**: `GET /livros`
   - **Descrição**: Retorna uma lista de todos os livros cadastrados.

### 2. Criar um novo livro
   - **Rota**: `POST /livros`
   - **Descrição**: Adiciona um novo livro ao banco de dados.
   - **Exemplo de Body JSON**:
     ```json
     {
       "titulo": "Livro 1",
       "autor": "Autor 1",
       "descricao": "Descrição do livro"
     }
     ```

### 3. Atualizar um livro
   - **Rota**: `PUT /livros/:id`
   - **Descrição**: Atualiza os dados de um livro específico.
   - **Exemplo de Body JSON**:
     ```json
     {
       "titulo": "Livro Atualizado",
       "autor": "Autor Atualizado",
       "descricao": "Nova descrição"
     }
     ```

### 4. Deletar um livro
   - **Rota**: `DELETE /livros/:id`
   - **Descrição**: Remove um livro específico do banco de dados.

## Testando com `curl`

Você pode testar os endpoints usando `curl` ou uma ferramenta como Postman.

#### 1. Listar todos os livros
```bash
curl http://localhost:8080/livros
```

#### 2. Criar um novo livro
```bash
curl -X POST http://localhost:8080/livros -d '{"titulo":"Livro Exemplo", "autor":"Autor Exemplo", "descricao":"Descrição exemplo"}' -H "Content-Type: application/json"
```

#### 3. Atualizar um livro
```bash
curl -X PUT http://localhost:8080/livros/1 -d '{"titulo":"Livro Atualizado", "autor":"Autor Atualizado", "descricao":"Nova descrição"}' -H "Content-Type: application/json"
```

#### 4. Deletar um livro
```bash
curl -X DELETE http://localhost:8080/livros/1
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou um pull request.

