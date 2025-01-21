# Stage 1: Build the Go application
FROM golang:1.23-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o main ./src/main.go

# Stage 2: Create a minimal production image
FROM alpine:latest

# Set the working directory inside the production container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/main .

# Run the application
CMD ["./main"]
