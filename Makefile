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
	$(GO_BINARY) test -race -coverprofile=coverage.txt -covermode=atomic $$( go list ./... | grep -v assert | grep -v future )

compare_files:
	@./compare_file.bash ./future/slices/slices_test.go ./future/slices/slices_1.22_test.go
	@./compare_file.bash ./future/maps/maps_test.go ./future/slices/maps_1.23_test.go
