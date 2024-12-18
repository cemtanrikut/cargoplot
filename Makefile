# Binary name
BINARY=cargoplot

# Default Go commands
GO=go
GOTEST=$(GO) test
GOBUILD=$(GO) build
GOMOD=$(GO) mod

# Default targets
all: build

build:
	$(GOBUILD) -o $(BINARY) ./cmd/main.go

run: build
	./$(BINARY)

test:
	$(GOTEST) ./...

coverage:
	$(GOTEST) -cover ./...

coverage-html:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

clean:
	rm -f $(BINARY) coverage.out

lint:
	golangci-lint run

# Phony targets (not actual files)
.PHONY: all build run test coverage coverage-html clean lint
