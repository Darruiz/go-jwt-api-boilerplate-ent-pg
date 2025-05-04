FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git gcc g++ libc-dev bash libpq

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates bash

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x ./main

COPY .env .env

EXPOSE 3009

CMD ["./main"]