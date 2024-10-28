# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PROJECT_DIR=$(shell pwd)
BIN_DIR=$(PROJECT_DIR)/bin
TEST_DIR=$(PROJECT_DIR)/test
BINARY_NAME=go-tut
BINARY_UNIX=$(BINARY_NAME)_unix
ENTRYPOINT=$(PROJECT_DIR)/cmd/go-tut/

# Colors
YELLOW=\033[33m
CYAN=\033[36m
RESET=\033[0m

.PHONY: all
all: test build

.PHONY: build
build:  ## Build the project
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(ENTRYPOINT)

.PHONY: run
run:  build ## Run the project
	$(BIN_DIR)/$(BINARY_NAME)

.PHONY: clean-test-cache
clean-test-cache:  ## Clean the test cache
	$(GOCLEAN) -testcache

.PHONY: test
test: clean-test-cache  ## Run the tests
	$(GOTEST) -v $(TEST_DIR)

.PHONY: clean
clean:  ## Clean the build files
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: deps
deps:  ## Install dependencies
	$(GOGET) -v $(PROJECT_DIR)

.PHONY: build-linux
build-linux:  ## Build the project for linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(ENTRYPOINT)

.PHONY: format
format:  ## Format the code
	$(GOCMD) fmt ./...

.PHONY: help
help:  ## Display this help screen
	@echo "Usage: make $(YELLOW)<target>$(RESET)"
	@echo "$(YELLOW)Available targets:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "$(CYAN)  %-15s$(RESET) - %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# make help default
.DEFAULT_GOAL := help
