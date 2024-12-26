# Use the official Go image
FROM golang:1.22

WORKDIR /app

# Copy Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o app .

# Expose the application port
EXPOSE 3003

# Run the application
CMD ["./app"]
