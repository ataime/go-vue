# Use the official golang image to create a build artifact.
FROM golang:1.16 as builder

# Set the current working directory inside the container.
WORKDIR /app

# Copy the Go modules files.
COPY src/go.mod src/go.sum ./

# Download and cache Go modules.
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container.
COPY src/main.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a minimal alpine image to run the Go app.
FROM alpine:latest

# Set the working directory to /app.
WORKDIR /app

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/app .

# Expose port 9090 to the outside world.
EXPOSE 9090

# Command to run the executable.
CMD ["./app"]
