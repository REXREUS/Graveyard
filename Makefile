.PHONY: build build-all clean run install

# Build for current platform
build:
	go build -o bin/graveyard cmd/graveyard/main.go

# Build optimized (smaller binary)
build-optimized:
	go build -ldflags="-s -w" -o bin/graveyard cmd/graveyard/main.go

# Build for all platforms
build-all:
	GOOS=windows GOARCH=amd64 go build -o bin/graveyard.exe cmd/graveyard/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/graveyard cmd/graveyard/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/graveyard-arm cmd/graveyard/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/graveyard-darwin cmd/graveyard/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/graveyard-darwin-arm cmd/graveyard/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Run without building
run:
	go run cmd/graveyard/main.go

# Install dependencies
install:
	go mod download
