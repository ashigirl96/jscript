REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X main.revision=$(REVISION)

## Install deps
.PHONY: deps
deps:
	go get -v -d
## Install dev-deps
.PHONY: dev-deps
dev-deps: deps
	go install github.com/Songmu/make2help/cmd/make2help@latest

# Run test
.PHONE: test
test: deps
	go test ./...

# Lint
.PHONY: lint
lint: dev-deps
	go vet ./...

# build binary
.PHONY: build
build: deps
	go build -ldflags "$(LDFLAGS)"

# SHOW help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)