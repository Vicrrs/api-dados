# api-dados

```
models/produto.go -> config/database.go -> controllers/produto_controller.go -> routes/produto_routes.go
```

```tree
api-dados/
├── main.go
├── models/
│   └── produto.go
├── controllers/
│   └── produto_controller.go
├── routes/
│   └── produto_routes.go
├── config/
│   └── database.go
└── go.mod
```


1. **Models**:  
   - Representam a estrutura dos dados da aplicação.  
   - São responsáveis por interagir com o banco de dados (consultas, inserções, atualizações, etc.).  
   - Exemplo: Um model `User` pode representar a tabela de usuários no banco de dados.

2. **Controllers**:  
   - Gerenciam a lógica da aplicação.  
   - Recebem as requisições, processam os dados (usando os models, se necessário) e retornam as respostas.  
   - Exemplo: Um controller `UserController` pode lidar com a criação, edição ou exclusão de um usuário.

3. **Routes**:  
   - Definem as URLs da aplicação e como elas são acessadas.  
   - Mapeiam as requisições HTTP (GET, POST, etc.) para os controllers correspondentes.  
   - Exemplo: Uma rota `/users` pode direcionar uma requisição GET para o método `index` do `UserController`.

Resumindo:  
- **Models** cuidam dos dados.  
- **Controllers** cuidam da lógica.  
- **Routes** cuidam das URLs e direcionam as requisições.  

