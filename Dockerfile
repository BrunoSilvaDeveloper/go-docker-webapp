# Use a imagem base do Go
FROM golang:1.20

# Crie e defina o diretório de trabalho
WORKDIR /app

# Copie o arquivo go.mod
COPY go.mod ./

# Baixe as dependências do módulo Go (se houver)
RUN go mod download

# Copie o código fonte para o container
COPY . .

# Compile a aplicação
RUN go build -o main .

# Defina o comando para iniciar a aplicação
CMD ["./main"]
