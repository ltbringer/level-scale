# test.Dockerfile → test image with Go installed
FROM golang:1.24-alpine

WORKDIR /app
RUN apk add --no-cache git
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest

ENV PATH="/root/go/bin:${PATH}"
COPY ../src/go.mod ./src/go.sum ./
RUN go mod download

COPY ../src ./
