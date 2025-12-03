.PHONY: build run stop docker-build docker-run

# Build the CLI tool
build:
	go build -o cve-search cmd/cve-search/main.go

# Kill any process on port 8081 and run the server
run: build
	@echo "Stopping any process on port 8081..."
	@-lsof -t -i:8081 | xargs -r kill -9
	@echo "Starting server..."
	go run cmd/server/main.go

# Docker commands
docker-build:
	docker build -t cve-backend .

docker-run:
	docker run -p 8081:8081 --env-file .env cve-backend

