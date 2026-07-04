.PHONY: help tidy test run-api run-travel run-agent clean build

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk -F ':.*?## ' '/^[a-zA-Z_-]+:.*?## / { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

tidy:	go mod tidy

test:	go test ./... -v -cover

run-api:	go run ./cmd/api/main.go

run-travel:	go run ./cmd/travel/main.go

run-agent:	go run ./cmd/agent/main.go

build:
	@echo "Building API Gateway..."
	go build -o bin/gateway ./cmd/api
	@echo "Building Travel Service..."
	go build -o bin/travel-service ./cmd/travel
	@echo "Building Agent Service..."
	go build -o bin/agent-service ./cmd/agent
	@echo "Build complete!"

clean:	rm -rf bin/ dist/
	go clean
