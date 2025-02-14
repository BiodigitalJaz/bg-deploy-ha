# Use the official Golang image as the build stage
FROM golang:1.22 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o api-server

# Use a minimal base image to run the compiled binary
FROM gcr.io/distroless/base-debian12

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=build /app/api-server .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./api-server"]
