
Este é um projeto de CRUD de livros desenvolvido em Golang, utilizando o framework **Gin** para manipulação de rotas e **MySQL** para armazenamento de dados. O projeto segue a arquitetura MVC (Model-View-Controller) para organização de código e está configurado para ser executado em AWS Lambda.

## Estrutura do Projeto

```bash
crud-livros
├── build/
│   └── main       # Arquivo Executavel Para Lambda
├── controllers/
│   └── routes.go       # Controladores das rotas de livros
│   └── livro_controller.go       # Controladores de livros
├── models/
│   └── livro.go                  # Modelo do Livro e funções CRUD
├── config/
│   └── db.go                     # Configuração do banco de dados
└── main.go              # Configuração do servidor e rotas principais                      
```

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
   go get -u github.com/awslabs/aws-lambda-go-api-proxy/gin
   ```

## Compilação do arquivo executavel para AWS Lambda

1. Compile o projeto para Linux:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o build/main main.go
   ```

2. Faça o upload do arquivo `main` para o AWS Lambda e configure o handler como `main`.

## Testando com `curl`

Você pode testar os endpoints usando `curl` ou uma ferramenta como Postman.

#### 1. Listar todos os livros
```bash
curl https://<your-api-gateway-endpoint>/livros
```

#### 2. Criar um novo livro
```bash
curl -X POST https://<your-api-gateway-endpoint>/livros -d '{"titulo":"Livro Exemplo", "autor":"Autor Exemplo", "descricao":"Descrição exemplo"}' -H "Content-Type: application/json"
```

#### 3. Atualizar um livro
```bash
curl -X PUT https://<your-api-gateway-endpoint>/livros/1 -d '{"titulo":"Livro Atualizado", "autor":"Autor Atualizado", "descricao":"Nova descrição"}' -H "Content-Type: application/json"
```

#### 4. Deletar um livro
```bash
curl -X DELETE https://<your-api-gateway-endpoint>/livros/1
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou um pull request.