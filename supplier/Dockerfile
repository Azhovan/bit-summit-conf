# Use a minimal base image for the final container
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o supplier

# Set the entry point for the container
ENTRYPOINT ["./supplier"]
