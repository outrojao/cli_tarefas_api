# Stage 1: Build
FROM golang:1.25.2-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código
COPY . .

# Compila a aplicação
RUN go build -o main ./cmd

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /root/

# Instala dependências se necessário (para conexões com banco, etc)
RUN apk --no-cache add ca-certificates

# Copia o binário compilado
COPY --from=builder /app/main .

# Copia as migrations se necessário
COPY --from=builder /app/internal/database/migrations ./internal/database/migrations/

# Copia configs se tiver
COPY --from=builder /app/configs ./configs/

EXPOSE 8000

CMD ["./main"]