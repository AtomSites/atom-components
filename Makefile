.PHONY: install generate test lint

install:
	@echo "Installing Go dependencies..."
	go mod download
	@echo "Installing development tools..."
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Installation complete!"

check-templ:
	@command -v templ >/dev/null 2>&1 || { echo "templ not found. Run: make install" >&2; exit 1; }

generate: check-templ
	@echo "Generating templ files..."
	templ generate

test: check-templ
	@echo "Running tests..."
	templ generate
	go test ./... -v

lint:
	@echo "Running linter..."
	golangci-lint run ./...
