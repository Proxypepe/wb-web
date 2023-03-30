FROM golang:1.18-alpine AS builder

WORKDIR /code

COPY . .

RUN go mod download

RUN go build -o build backend/main.go

FROM alpine:3.14

WORKDIR /code

COPY --from=builder /code/build .

EXPOSE 8080

ENV POSTGRES_DB=wb \
    POSTGRES_HOST=postgres \
    POSTGRES_PASSWORD=postgres \
    POSTGRES_PORT=5432 \
    POSTGRES_USER=alex \
    REDIS_HOST=redis \
    REDIS_PORT=6379 \
    NATS_CLUSTER_ID=test-cluster\
    NATS_CLIENT_ID=test-client\
    NATS_HOST=nats\
    NATS_PORT=4222\
    NATS_SUBJECT=backend\
    SERVER_HOST=0.0.0.0 \
    SERVER_PORT=8080

CMD ["./build"]