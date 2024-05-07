# Use the official Go image as a base
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy the client code into the container
COPY . .

# Download all the dependencies
RUN go mod tidy

# Compile the Go app
RUN go build -o client

# Command to run the executable
CMD ["./client"]
