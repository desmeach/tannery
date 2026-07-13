BINARY_NAME=tannery
BUILD_DIR=./bin
MAIN_PATH=./cmd/app/main.go

GOOS?=linux
GOARCH?=amd64

RM=rm -rf

.PHONY: all build clean run test lint fmt security deps help

all: fmt lint test build

## clean: Build application
build:
	@echo "==> Building app..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

## clean: Cross build application
build-cross:
	@echo "==> Cross compile for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	ENV GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH) $(MAIN_PATH)

## clean: Clean application build
clean:
	@echo "==> Cleaning app..."
	@$(RM) $(BUILD_DIR)
	@go clean -testcache
	@echo "==> Done"

## run: Run application
run:
	@echo "==> Run app..."
	@exec go run $(MAIN_PATH) --config_path=$(CONFIG_PATH)

## test: Run all tests
test:
	@echo "==> Run tests..."
	go test -v -race -cover ./...

## fmt: Format project code
fmt:
	@echo "==> Format code..."
	go fmt ./...

## lint: Linter for project code
lint:
	@echo "==> Code lint..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "Error: golangci-lint not installed. Download it from https://golangci-lint.run"; \
		exit 1; \
	fi

## deps: Install all project dependencies
deps:
	@echo "==> Process all dependencies..."
	go mod download
	go mod tidy

## docker-build: Build app docker image
docker-build:
	@echo "==> Build docker image..."
	docker build -t $(BINARY_NAME):latest .

## help: Show available make commands
help:
	@echo "Commands:"
	@grep -E '^## .*' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ": "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' | sed 's/## //'

## swag: Run swagger build
swag:
	@echo "Generating swagger docs..."
	swag i -g cmd/app/main.go -o cmd/docs -d . --parseDependency --parseInternal