.phony: test check lint

GO_BINARY ?= go

check: lint
	@$(GO_BINARY) vet ./...
	@gofmt -l .
	@test -z "$$( gofmt -l . )"

lint:
ifndef SKIP_LINT
	@golangci-lint run ./...
endif

test:
	$(GO_BINARY) test -race -v ./...

cov: SHELL:=/bin/bash
cov:
	$(GO_BINARY) test -race -coverprofile=coverage.txt -covermode=atomic $$( $(GO_BINARY) list ./... | grep -v assert )
