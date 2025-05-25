FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o lamp-shop cmd/main.go

FROM alpine:3.17
WORKDIR /app

COPY --from=builder /app/lamp-shop .
COPY --from=builder /app/migrations ./migrations

RUN apk add --no-cache bash postgresql-client

ENTRYPOINT ["./lamp_backend"]