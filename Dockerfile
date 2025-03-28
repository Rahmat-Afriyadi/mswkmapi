# Stage 1: Build the application
FROM golang:1.22

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o app .

# Stage 2: Create a lightweight container for the application
# FROM debian:bookworm-slim

# WORKDIR /app

# Copy only the built binary from the builder stage
# COPY --from=builder /app/app .

# Expose the application port
EXPOSE 3003

# Run the application
CMD ["./app"]
