# Build Stage
FROM golang:1.24-alpine AS builder

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

# Install go-exploitdb
RUN go install github.com/vulsio/go-exploitdb@latest

# Fetch ExploitDB database during build
RUN /go/bin/go-exploitdb fetch exploitdb --dbpath /app/go-exploitdb.sqlite3
RUN /go/bin/go-exploitdb fetch awesomepoc --dbpath /app/go-exploitdb.sqlite3
RUN /go/bin/go-exploitdb fetch githubrepos --dbpath /app/go-exploitdb.sqlite3
RUN /go/bin/go-exploitdb fetch inthewild --dbpath /app/go-exploitdb.sqlite3

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies (if any)
RUN apk add --no-cache ca-certificates

# Copy binaries from builder
COPY --from=builder /app/cve-search .
COPY --from=builder /app/server .
COPY --from=builder /go/bin/go-exploitdb /usr/local/bin/go-exploitdb

# Copy ExploitDB database
RUN mkdir -p /root/.go-exploitdb
COPY --from=builder /app/go-exploitdb.sqlite3 /root/.go-exploitdb/go-exploitdb.sqlite3

# Expose the server port
EXPOSE 8081

# Run the server
CMD ["./server"]
