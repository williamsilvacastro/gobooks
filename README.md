
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


## TO DO

1.  Melhorar resiliencia 

2.  Variaveis de ambiente para conexao com banco

3.  Validacao se tabela de livros existe no banco

4.  Criar banco RDS para este projeto

5.  Melhorar tratamento de erro e response http status que estamos usando

6.  Fazer um CICD Simplificado


## Sugestao GITHUB Copilot:

1.  Tratamento de Erros e Respostas HTTP:
Melhore o tratamento de erros e as respostas HTTP para fornecer mensagens mais detalhadas e específicas.
Utilize middlewares para capturar e tratar erros de forma centralizada.

2.  Validação de Dados:
Adicione validações mais robustas para os dados de entrada usando bibliotecas como go-playground/validator.

3.  Autenticação e Autorização:
Implemente autenticação e autorização para proteger as rotas da API.
Considere o uso de JWT (JSON Web Tokens) para autenticação.

4.  Documentação da API:
Utilize ferramentas como Swagger para documentar a API.
Forneça exemplos de requisições e respostas para cada endpoint.

5.  Testes Automatizados:
Aumente a cobertura de testes unitários e de integração.
Utilize mocks para testar interações com o banco de dados e outras dependências externas.

6.  Logs e Monitoramento:
Adicione logs detalhados para facilitar a depuração e o monitoramento.
Utilize ferramentas de monitoramento e alertas para acompanhar a saúde da aplicação.

7.  Configurações e Variáveis de Ambiente:
Utilize variáveis de ambiente para configurar a aplicação, especialmente para informações sensíveis como credenciais de banco de dados.
Considere o uso de bibliotecas como viper para gerenciar configurações.

8.  CICD (Continuous Integration and Continuous Deployment):
Configure pipelines de CI/CD para automatizar testes, builds e deploys.
Utilize ferramentas como GitHub Actions, Jenkins ou GitLab CI.

9.  Segurança:
Realize auditorias de segurança para identificar e corrigir vulnerabilidades.
Utilize práticas recomendadas de segurança, como a sanitização de entradas e a proteção contra ataques de injeção SQL.

10.  Escalabilidade:
Considere a escalabilidade da aplicação, especialmente se for implantada em ambientes de produção.
Utilize práticas de design de software que suportem a escalabilidade horizontal e vertical.

11.  Boas Práticas de Código:
Siga as convenções de código do Go, como as recomendadas pelo golangci-lint.
Realize revisões de código para garantir a qualidade e a consistência.