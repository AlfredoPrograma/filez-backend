# Builds application binary
build:
	go build -o filezserver ./...

# Spins up development mode with hot reload
dev:
	air -c ./.air.toml

# Run tests
test:
	go test -v ./...
