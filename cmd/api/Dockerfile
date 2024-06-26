# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16 as builder

# Create and change to the app directory.
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
FROM alpine:latest

# Copy the binary to the production image from the builder stage.
COPY --from=builder /bin/api /api

# Run the web service on container startup.
CMD ["/api"]
