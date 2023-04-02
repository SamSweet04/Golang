# Use an official Golang runtime as a parent image
FROM golang:1.17.4-alpine3.14

# Set the current working directory inside the container
WORKDIR /app

# Copy the dependencies files to the container
COPY go.mod ./
COPY go.sum ./

# Install the dependencies
RUN go mod download

# Copy the rest of the application to the container
COPY . .

# Build the application
RUN go build -o bookstore .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application
CMD ["./bookstore"]
