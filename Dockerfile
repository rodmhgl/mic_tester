FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Create minimal image 
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy built application from previous stage
COPY --from=builder /app/app .

# Set environment variables
ENV PORT=8080
ENV APP_ENV=production

# Expose port
EXPOSE 8080

# Run the application
ENTRYPOINT ["./app"]