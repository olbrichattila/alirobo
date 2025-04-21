# Use official Golang image as builder
FROM golang:1.24 AS builder

# Set working directory inside the container
WORKDIR /app

RUN apt install gcc \
    && apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config

# Copy go.mod and go.sum
COPY * ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go server
RUN go build -o server .

# Final lightweight image
FROM debian:bookworm-slim

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose the application port
EXPOSE 8080

# Run the server binary
CMD ["./server"]
