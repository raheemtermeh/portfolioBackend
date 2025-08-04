# Dockerfile for Go Backend

# --- Build Stage ---
# Use the official Go image to build the application
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
COPY . .
# Build the Go application
# CGO_ENABLED=0 is important for creating a static binary
# -o /app/main creates the output file named 'main' in the /app directory
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# --- Production Stage ---
# Use a minimal image for the final container
FROM alpine:latest
WORKDIR /app
# Copy the built binary from the builder stage
COPY --from=builder /app/main .
# Copy the .env file (we will create this on the server)
COPY .env .
# Expose the port the Go application runs on
EXPOSE 8080
# Command to run the executable
CMD ["./main"]