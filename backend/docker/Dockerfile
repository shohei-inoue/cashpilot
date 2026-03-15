# ========== 本番用ビルド（デフォルト） ==========
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/api

FROM alpine:3.21 AS prod
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

# ========== 開発用（Air ホットリロード） ==========
FROM golang:1.25-alpine AS dev
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080
CMD ["air", "-c", "docker/.air.toml"]
