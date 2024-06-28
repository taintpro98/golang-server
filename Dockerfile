# Build stage
FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Building app
RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux go build -o api

# Final stage
FROM alpine:3.18

# Tạo thư mục config nếu chưa tồn tại
RUN mkdir -p /app/config

COPY --from=builder /app/cmd/api/api /app/
COPY ./config/config.mn.yml /app/config/config.yml

ENV ENV_CONFIG_ONLY=true

WORKDIR /app

EXPOSE 8080
ENV GIN_MODE=release

# Run the web service on container startup.
CMD ["/app/api"]
