FROM golang:1.21.5-alpine AS build

WORKDIR /app

# Copy only the necessary Go application files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application for linux
# CGO_ENABLED=0 since we don't have external dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Create a non-root user
RUN adduser --no-create-home --disabled-password -u 10001 --shell /bin/sh appuser && \
    chown -R appuser /app

# Set the build cache directory to a location where the user has write permissions
ENV GOCACHE=/tmp/.cache/go-build

USER appuser

# Expose the port for swagger
EXPOSE 1323

# Run the Go application when the container starts
CMD ["./main"]