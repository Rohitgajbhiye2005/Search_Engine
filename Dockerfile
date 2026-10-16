# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o mini-search-engine .

# Stage 2: Run the binary in a tiny alpine container
FROM alpine:latest

WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/mini-search-engine .

# COPY your frontend/static assets folder so the app can serve them!
# (Change 'static' to whatever your folder is named, e.g., 'templates', 'frontend', etc.)
COPY --from=builder /app/frontend ./frontend

EXPOSE 8080

CMD ["./mini-search-engine"]