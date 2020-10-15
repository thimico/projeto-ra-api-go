# Documentação Backend Engineer Challenge
  Projeto Api Challenge RA
  
### Ferramentas ###
  
  * golang >= 13 (test&development)
  
  * mongodb >= 2.6.10 (test&development)
  
  * docker >= 17.03.2-ce (development&production)
  
  * redis >= 3.0.6 (development&production)
  
  * nginx >= 1.0.15 (production)
  
# Organização dos Diretórios:
Pasta Principal é /cmd , lá contém o arquivo (main.go) que é o arquivo para executar as chamadas de metódos já com o teste de conexão. 

-> Diretório pkg/api/model

  contém o arquivo complain.go que contem as structs utilizadas no projeto  

# Funcionamento dos Arquivos
# Caminho : /router/routes.go

O arquivo routes tem acesso as funções correspondentes aos paths complains, health e home

O código é compativel para testagem via Postman

# Caminho : /util/mg/mongo.go
imports de fora do projeto:

"go.mongodb.org/mongo-driver/mongo" -> para conectar com o banco

# Caminho: controller/complainController.go
Funções 

SaveComplain()-> Nessa função vc cria uma Reclamação

FindByParam() -> Busca todas as Reclamações baseados nos parametros opcionalmente

FindById() -> Pesquisa um registro por seu ID

Update()-> Atualiza um registro

DeleteById() -> deleta o registro por seu ID 

# Executando com docker-compose
Vá até a pasta do projeto e rode o seguinte comando

`docker-compose up`

for production: `docker-compose -f docker-compose-prod.yml up`

# Executando manualmente

Requistos : Necessário criar um server/banco de dados mongo

Vá até a pasta do projeto e rode os seguintes comandos

`go build ./cmd/main.go` para buildar e baixar todas as dependências

`./main`  para executar o projeto após o build

# Testes

para rodar os testes bastar ir até a pasta correspondente ou raiz e rodar o seguinte comando
`go test ./...`

Também pode ser feito testes atráves da ferramenta postman:

http://localhost:8080/complains
POST ->
	Body:
```json
{
  "title": "Nenhuma atenção com o cliente",
  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
  "locale": {
    "city": "Savador",
    "state": "Bahia"
  },
  "company": {
    "title": "Claro",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
  }
}
```

http://localhost:8080/complains
GET ->
	Body:
```json
[
  {
     "id": "5f88529896a156557ef2a813",
     "title": "Nenhuma atenção com o cliente",
     "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
     "locale": {
       "city": "Rio de Janeiro",
       "state": "Rio de Janeiro"
     },
     "company": {
       "title": "Instagram",
       "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
     },
     "countPageViews": 4161611,
     "isOnTop10BadRA": true
   },
   {
     "id": "5f8852c196a156557ef2a816",
     "title": "Nenhuma atenção com o cliente",
     "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
     "locale": {
        "city": "Savador",
        "state": "Bahia"
     },
     "company": {
       "title": "Claro",
       "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
     },
     "countPageViews": 1634763,
     "isOnTop10BadRA": true
   }
 ]
```

http://localhost:8080/complains?company=Claro
GET ->
	Body:
```json
[
  {
      "id": "5f8852c196a156557ef2a816",
      "title": "Nenhuma atenção com o cliente",
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      "locale": {
         "city": "Savador",
         "state": "Bahia"
      },
      "company": {
        "title": "Claro",
        "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
      },
      "countPageViews": 1634763,
      "isOnTop10BadRA": true
  }
]
```

http://localhost:8080/complains?city=Rio%20de%20Janeiro
GET ->
	Body:
```json
[
  {
       "id": "5f88529896a156557ef2a813",
       "title": "Nenhuma atenção com o cliente",
       "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
       "locale": {
         "city": "Rio de Janeiro",
         "state": "Rio de Janeiro"
       },
       "company": {
         "title": "Instagram",
         "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
       },
       "countPageViews": 4161611,
       "isOnTop10BadRA": true
  }
]
```

http://localhost:8080/complains/5f88529896a156557ef2a813
GET ->
	Body:
```json
{
   "id": "5f88529896a156557ef2a813",
   "title": "Nenhuma atenção com o cliente",
   "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
   "locale": {
     "city": "Rio de Janeiro",
     "state": "Rio de Janeiro"
   },
   "company": {
     "title": "Instagram",
     "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
   },
   "countPageViews": 4161611,
   "isOnTop10BadRA": true
}
```

http://localhost:8080/complains/5f88529896a156557ef2a813
PUT ->
	Body:
```json
    {
      "title": "Nenhuma atenção com o cliente",
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      "locale": {
        "city": "Rio de Janeiro",
        "state": "Rio de Janeiro"
      },
      "company": {
        "title": "Magazine Luiza",
        "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
      }
    }
```

http://localhost:8080/complains/5f88529896a156557ef2a813
DELETE ->
	

# Especificações
go version 14 linux/amd64

Mongo

Sistema Operacional: linux Ubuntu 18.4

# Importações
`go build ./...`

### Swagger Documentation ###

  
Swagger disponibilizado em: `{{URL}}:{{PORT}}/swagger/index.html`

# Contatos
thg.mnzs@gmail.com
