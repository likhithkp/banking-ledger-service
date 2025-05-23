FROM golang:1.24.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN apt-get update && apt-get install -y librdkafka-dev

RUN go build -o app main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates curl librdkafka1 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/db/psql/migrations ./db/psql/migrations

EXPOSE 3001
CMD ["./app"]
