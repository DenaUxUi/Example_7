# Этап 1: билдим бинарник
FROM golang:1.22.2 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o server

# Этап 2: минимальный контейнер
FROM debian:bookworm-slim

# Переменные окружения (если хочешь задать порт)
ENV PORT=8081

# Копируем бинарник из билдера
COPY --from=builder /app/server /server

# Указываем порт (документально, не обязательно)
EXPOSE 8081

# Запуск сервера
ENTRYPOINT ["/server"]
