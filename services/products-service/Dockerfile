FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/products-service ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/products-service .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/seeds ./seeds
COPY --from=builder /app/.env .env

EXPOSE 8080

ENTRYPOINT ["./products-service"]