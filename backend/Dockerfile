FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
COPY migrations ./migrations

RUN go mod tidy
RUN go build -o server ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server ./server
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./server"]