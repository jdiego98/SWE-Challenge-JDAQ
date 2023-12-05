# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the Go mod and sum files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Run the application
CMD ["./main"]
