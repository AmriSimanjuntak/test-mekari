FROM golang:alpine as builder

# Set the current working directory inside the container
WORKDIR /amri/go/src/test-mekari/

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies listed in go.mod and cache them
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o test-mekari

# Use a lightweight Alpine image to run the application
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /amri/go/src/test-mekari/

# Copy the built executable from the builder stage to the working directory
COPY --from=builder /amri/go/src/test-mekari .

COPY --from=builder /amri/go/src/test-mekari/docker-compose.yml .

ENV TZ=Asia/Jakarta


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./test-mekari"]
