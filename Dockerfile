# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files first for efficient caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o mini-search-engine .

# Stage 2: Run the binary in a tiny alpine container
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/mini-search-engine .

# Expose port 8080 for your Gin REST API
EXPOSE 8080

# Command to run your application
CMD ["./mini-search-engine"]