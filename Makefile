# Variables
BINARY_NAME := finalize-cli
INPUT_FILE ?= inputfiles/financial_report.csv
DATE_RANGE ?= "2024-01-01:2024-02-01"
COVERAGE_FILE := test/coverage.out
COVERAGE_HTML := test/coverage.html

# Default target: Build the project
build:
	@echo "Building the project..."
	go build -o bin/$(BINARY_NAME) cmd/main.go

# Run the application with flags
run: build
	@echo "Running the application..."
	./bin/$(BINARY_NAME) -input $(INPUT_FILE) -date $(DATE_RANGE)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -f bin/$(BINARY_NAME) $(COVERAGE_FILE) $(COVERAGE_HTML)

# Run tests with verbose output
test:
	@echo "Running all tests..."
	go test ./... -v

# Run tests with code coverage
coverage:
	@echo "Running tests with coverage..."
	go test ./... -v -coverprofile=$(COVERAGE_FILE)
	@echo "Generating HTML coverage report..."
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"

# Open coverage report in the default browser (macOS/Linux)
open-coverage: coverage
	@echo "Opening coverage report..."
	@xdg-open $(COVERAGE_HTML) 2>/dev/null || open $(COVERAGE_HTML) 2>/dev/null || echo "Please open $(COVERAGE_HTML) manually."

# Run tests with the race detector
race:
	@echo "Running tests with the race detector..."
	go test ./... -race -v

# Lint the code (requires golangci-lint to be installed)
lint:
	@echo "Linting the code..."
	golangci-lint run

# Format the code
fmt:
	@echo "Formatting the code..."
	go fmt ./...

# Run all checks (lint, fmt, test)
all: fmt lint test
	@echo "All checks passed!"
