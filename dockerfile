# Используем официальный образ Golang
FROM golang:1.20

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию контейнера
COPY . .

# Переменные окружения для подключения к PostgreSQL
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=postgres

# Команда для запуска приложения
CMD ["go", "run", "./cmd/web"]
