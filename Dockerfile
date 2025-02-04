# Use latest Go version to avoid toolchain directive errors
FROM golang:1.23 AS builder

# Set environment variable for Go modules
ENV GO111MODULE=on

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory in the new container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Copy the .env file (optional)
COPY .env .

# Install bash and libc6-compat (optional for debugging and dependencies)
RUN apk add --no-cache bash libc6-compat

# Run the application
CMD ["./main"]