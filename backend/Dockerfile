# Build stage
FROM golang:1.25.10-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Run stage
FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]