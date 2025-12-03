# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod and sum files
COPY go.mod ./
# COPY go.sum ./ # Uncomment if go.sum exists

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the CLI tool and Server
# We build both because the current architecture requires the server to exec the CLI binary
RUN go build -o cve-search cmd/cve-search/main.go
RUN go build -o server cmd/server/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies (if any)
RUN apk add --no-cache ca-certificates

# Copy binaries from builder
COPY --from=builder /app/cve-search .
COPY --from=builder /app/server .

# Expose the server port
EXPOSE 8081

# Run the server
CMD ["./server"]
