FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/orders-service ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/bin/orders-service .
COPY .env .env

EXPOSE 8090

ENTRYPOINT ["./orders-service"]