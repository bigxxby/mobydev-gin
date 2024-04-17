FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

COPY . .

CMD ["go", "run", "./app/"]
