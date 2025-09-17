# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
BINARY_NAME=booru-server
BINARY_UNIX=$(BINARY_NAME)_unix

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/server

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/server
	./$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: all build run test clean
