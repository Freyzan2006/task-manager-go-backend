FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o task-manager ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/task-manager .
COPY migrations ./migrations
CMD ["./task-manager"]
