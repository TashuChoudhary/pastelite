FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder app/server .
COPY public/ ./public/

EXPOSE 3000

CMD ["./server"]