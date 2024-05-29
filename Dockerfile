# Start from the official Golang image with version 1.22.3
FROM golang:1.22.3 AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files first
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
