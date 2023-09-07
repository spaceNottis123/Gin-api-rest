# To start this project
1 - docker compose up
2 - go run main.go 


# API Calls

POST
curl --location 'http://localhost:8080/alunos/' \
--header 'Content-Type: application/json' \
--data '{
    "nome":"Gui Ramos",
    "rg":"1234563258",
    "cpf":"40255111230"
}'

GET - todos alunos
curl --location 'http://localhost:8080/alunos/'

GET - aluno por ID
curl --location 'http://localhost:8080/alunos/1'

GET - aluno por CPF
curl --location 'http://localhost:8080/alunos/cpf/40255111230'

DELETE - deleta aluno por ID
curl --location 'http://localhost:8080/alunos/cpf/40255111230'

PATCH - atualiza aluno por id + json
curl --location --request PATCH 'http://localhost:8080/alunos/4' \
--header 'Content-Type: application/json' \
--data '{
    "nome":"Gui Ramos Juninho",
    "rg":"1234563258",
    "cpf":"40255111230"
}'