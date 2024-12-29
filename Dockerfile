############################
# Stage 1: Builder
############################
FROM golang:1.23-alpine3.19 AS builder

# Create and set working directory
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
# The output binary is placed in /app/bin/server
RUN mkdir -p bin && \
    go build -o bin/server ./cmd/server

############################
# Stage 2: Minimal Runtime
############################
FROM alpine:latest

# Set environment variables
ENV GIN_MODE=release

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/server /app/server

# Copy configuration files
COPY --from=builder /app/config/ /app/config/

# Expose port (ensure it matches your server's port)
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
