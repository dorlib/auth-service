# Use Golang base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

RUN go env -w GOPROXY=https://goproxy.io,direct

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o auth-service .

# Expose port 8081
EXPOSE 8086

# Command to run the executable
CMD ["./auth-service"]
