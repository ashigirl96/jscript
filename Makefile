REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X main.revision=$(REVISION)

## Install deps
.PHONY: deps
deps:
	go mod download
## Install dev-deps
.PHONY: dev-deps
dev-deps: deps
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/Songmu/make2help/cmd/make2help@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

# Run test
.PHONE: test
test: deps
	go test -race ./...

# Lint
.PHONY: lint
lint: dev-deps
	golangci-lint run ./...
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/your/package"
	git mod tidy

# build binary
.PHONY: build
build: deps
	go build -ldflags "$(LDFLAGS)"

# SHOW help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)