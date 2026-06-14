# Stage 1: Builder
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -trimpath \
    -o /server ./cmd/app/main.go

# Stage 2: Production
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /server /server

COPY --from=builder /app/configs /configs

# Порт
EXPOSE 8080

# Запуск
CMD ["/server"]