# Use a recent Go base image (consider upgrading if necessary)
FROM golang:1.22

# Set working directory
WORKDIR /app

# Download Go modules (optimize by caching)
COPY go.mod go.sum ./
RUN go env
RUN go mod download 

# Copy the source code
COPY *.go ./

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional: Expose port for TCP communication
EXPOSE 4056

# Run the application
CMD ["/docker-gs-ping"]
