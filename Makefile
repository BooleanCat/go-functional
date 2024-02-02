.phony: test check lint

GO_122 = go

check: lint
	@go vet ./...
	@gofmt -l .
	@test -z "$$( gofmt -l . )"

lint:
ifndef SKIP_LINT
	@golangci-lint run ./...
endif

test:
	go test -race -v ./...

test-v2: export GOEXPERIMENT = rangefunc
test-v2:
	@cd v2; $(GO_122) vet ./...
	@cd v2; $(GO_122) test -v -race ./...

cov:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
