# Stage 1: Build the Go application with CGO enabled
FROM golang:1.23 AS builder

# Enable CGO and set up environment variables for Go
ENV CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 CC=arm-linux-gnueabihf-gcc

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code, including configs directories
COPY . .

# Install the necessary development libraries for SQLite and cross-compilation tools
RUN apt-get update && apt-get install -y gcc libc6-dev gcc-arm-linux-gnueabihf

# Build the application
RUN go build -o database-handler ./main.go

# Stage 2: Create a lightweight runtime environment with Debian
FROM debian:bookworm-slim

# Install necessary runtime libraries and CA certificates
RUN apt-get update && apt-get install -y libc6 libsqlite3-0 ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the built binary and directories from the builder stage
COPY --from=builder /app/database-handler /database-handler
COPY --from=builder /app/configs /configs
COPY --from=builder /app/sql /sql
COPY --from=builder /app/secret-data /secret-data

# Default command keeps container alive
CMD ["tail", "-f", "/dev/null"]
