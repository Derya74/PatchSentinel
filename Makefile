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
