# Builder
FROM golang:1.24-alpine AS builder

WORKDIR /capibara

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

# Runtime
FROM alpine:latest

WORKDIR /capibara

COPY --from=builder /capibara/server .