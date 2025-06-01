# BUILD
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY ../src/go.mod ./src/go.sum ./
RUN go mod download

COPY ../src ./src
WORKDIR /app/src

# Build the binary (statically linked, no CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# OPTIMIZE
FROM alpine:3.19

# Install CA certificates (required for HTTPS calls)
RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=builder /app/src/app .
