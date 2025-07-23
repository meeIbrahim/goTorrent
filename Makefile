# Go parameters
VERSION ?= $(shell git describe --tags --always --dirty)
BINARY_NAME=goTorrent
BUILD_DIR=build


.PHONY: all build clean run

all: build


build: clean
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION) -ldflags "-X main.version=$(VERSION)"

run: build
	@echo "Running..."
	@./$(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

