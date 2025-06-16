# Production Dockerfile
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a smaller image for the final stage
FROM alpine:latest

WORKDIR /app

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/problems /app/problems
COPY --from=builder /app/code_templates /app/code_templates

# Expose port 5000
EXPOSE 5000

# Run the application
CMD ["./main"]
