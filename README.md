# Meu Projeto Go com Docker

Este projeto é uma aplicação web simples escrita em Go, que responde a requisições HTTP. A aplicação foi empacotada em um container Docker utilizando Docker Compose.

## Funcionalidades

- **GET /**: Retorna uma resposta com o texto "Olá, Mundo!".
- **GET /status**: Retorna um JSON com o status da aplicação, por exemplo: `{ "status": "ok" }`.

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em sua máquina:

- [Go](https://golang.org/doc/install) (versão 1.20 ou superior)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Configuração do Ambiente

Este projeto permite que você defina o ambiente de execução utilizando uma variável de ambiente `APP_ENV`. No `docker-compose.yml`, você pode definir essa variável como `development`, `production`, ou qualquer outro valor.

## Como Executar o Projeto

Para executar o projeto, siga os passos abaixo:

1. **Clonar o Repositório**: Clone o repositório para sua máquina local:
    ```sh
    git clone https://github.com/seu-usuario/seu-repositorio.git
    cd seu-repositorio
    ```

2. **Compilar e Executar a Aplicação Localmente**: Para compilar e executar a aplicação localmente (sem Docker):
    ```sh
    go mod tidy
    go run main.go
    ```
    A aplicação estará disponível em [http://localhost:8080](http://localhost:8080).

3. **Executar a Aplicação com Docker**:
    - **Construir a Imagem Docker**: Para construir a imagem Docker da aplicação, execute:
        ```sh
        docker-compose build
        ```
    - **Iniciar o Container**: Para iniciar o container com a aplicação, execute:
        ```sh
        docker-compose up
        ```
    A aplicação estará disponível em [http://localhost:8080](http://localhost:8080).

## Estrutura do Projeto

A estrutura do projeto é a seguinte:

```plaintext
.
├── Dockerfile
├── docker-compose.yml
├── main.go
├── go.mod
└── README.md
