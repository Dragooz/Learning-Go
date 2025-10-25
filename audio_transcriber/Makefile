# Chinese Audio Transcriber Makefile

.PHONY: build run clean test deps

# Build the application
build:
	go build -o bin/transcriber ./cmd/transcriber

# Run the application
run: build
	./bin/transcriber

# Run with custom config
run-config: build
	./bin/transcriber -config=./config.yaml

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Run tests
test:
	go test ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Install dependencies
install-deps:
	go mod download

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Build for Windows (cross-compile)
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/transcriber.exe ./cmd/transcriber

# Development build with race detection
build-dev:
	go build -race -o bin/transcriber-dev ./cmd/transcriber

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Build and run the application"
	@echo "  run-config   - Build and run with custom config file"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  lint         - Run linter"
	@echo "  build-windows- Cross-compile for Windows"
	@echo "  build-dev    - Build with race detection"
	@echo "  help         - Show this help"