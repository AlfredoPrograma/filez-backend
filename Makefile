# Builds application binary
build:
	go build -o filezserver ./cmd/api/main.go

# Spins up development mode with hot reload
dev:
	air -c ./.air.toml

# Run tests
tests:
	go test -v ./...
