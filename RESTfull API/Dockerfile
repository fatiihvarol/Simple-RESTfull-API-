# Base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port that the application will listen on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
