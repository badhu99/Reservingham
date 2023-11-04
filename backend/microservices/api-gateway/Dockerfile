# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . /app

# Build the Go application
RUN go build -o main

# Expose the port specified in the environment variable
EXPOSE $PORT

# Run the Go application
CMD ["/app/main"]