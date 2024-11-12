# Variables
BINARY_NAME := finalize-cli
INPUT_FILE ?= inputfiles/financial_report.csv
DATE_RANGE ?= "2024-01-01:2024-02-01"

# Default target: Build the project
build:
	@echo "Building the project..."
	go build -o bin/$(BINARY_NAME) cmd/main.go

# Run the built binary with flags
run: build
	@echo "Running the application..."
	./bin/$(BINARY_NAME) -input $(INPUT_FILE) -date $(DATE_RANGE)

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	rm -f ./bin/$(BINARY_NAME)

# Cross-compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux cmd/main.go
