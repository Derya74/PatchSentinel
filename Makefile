# PatchSentinel — build, test and demo automation.
#
# Requires: Go >= 1.24, Node >= 20. No other tools are needed.

GO       ?= go
NODE     ?= node
NPM      ?= npm
BINDIR   := bin
BIN      := $(BINDIR)/patchsentinel
PKG      := ./...
VERSION  ?= 1.0.0
LDFLAGS  := -X main.version=$(VERSION)

.PHONY: all build test test-go test-ts vet fmt cover clean demo ci web-install help

all: build test ## Build everything and run all tests

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

build: ## Compile the Go CLI into ./bin
	@mkdir -p $(BINDIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/patchsentinel

vet: ## Run go vet
	$(GO) vet $(PKG)

fmt: ## Format Go sources
	$(GO) fmt $(PKG)

test-go: ## Run Go tests
	$(GO) test $(PKG)

cover: ## Run Go tests with coverage report
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -func=coverage.out

web-install: ## Install TypeScript dev dependencies
	cd web && $(NPM) install --no-audit --no-fund

test-ts: web-install ## Type-check, build and test the TypeScript renderer
	cd web && $(NPM) test
