# Этап 1: Сборка frontend
FROM node:18-bullseye-slim as frontend-builder

WORKDIR /app/frontend

COPY frontend/package.json frontend/package-lock.json ./

RUN npm install

COPY frontend ./

RUN npm run build

RUN rm -rf /app/frontend

# Этап 2: Итоговый контейнер
FROM golang:1.21-bullseye

WORKDIR /app

# Копируем файлы go.mod и go.sum для загрузки зависимостей Go
COPY go.mod go.sum ./

# Копируем статические файлы из этапа сборки frontend
COPY --from=frontend-builder /app/static /app/static

# Копируем файлы backend
COPY backend /app/backend

# Загружаем зависимости Go
RUN go mod download

# Собираем бинарный файл
RUN go build -o main /app/backend/main.go

# Открываем порт, на котором будет работать приложение
EXPOSE 8082

# Запускаем приложение
CMD ["/app/main"]
