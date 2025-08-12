FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]
