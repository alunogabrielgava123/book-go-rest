# API de CRUD de Books

Essa é uma aplicação API REST desenvolvida em Go utilizando o framework Gin para realizar operações CRUD na tabela de livros.

## Endpoints

### GET /books

Retorna uma lista de todos os livros cadastrados.

para solicitar esse tipo re requisição é necessario ter um berer token.

Exemplo de resposta:
[
{
"id": 1,
"title": "O Pequeno Príncipe",
"author": "Antoine de Saint-Exupéry",
"price": 29.9
},
{
"id": 2,
"title": "Dom Quixote",
"author": "Miguel de Cervantes",
"price": 34.9
}
]

### GET /books/:id

Retorna um livro específico pelo ID.

Exemplo de resposta:
{
"id": 1,
"title": "O Pequeno Príncipe",
"author": "Antoine de Saint-Exupéry",
"price": 29.9
}

### POST /books

Cria um novo livro.

Exemplo de corpo da requisição:
{
"title": "A Revolução dos Bichos",
"author": "George Orwell",
"price": 24.9
}

### PUT /books/:id

Atualiza um livro existente pelo ID.

Exemplo de corpo da requisição:
{
"title": "A Revolução dos Bichos",
"author": "George Orwell",
"price": 29.9
}

### DELETE /books/:id

Exclui um livro existente pelo ID.

## Instalação e execução

1. Clone o repositório: `git clone https://github.com/seu-usuario/projeto-exemplo.git`.
2. Navegue até a pasta raiz do projeto e execute `go run main.go`.
3. Acesse a API em `http://localhost:8080`.