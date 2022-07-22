.phony: test check

check:
	@go vet ./...
	@gofmt -l .
	@test -z "$$( gofmt -l . )"
	@golangci-lint run ./...

test:
	go test -v ./...
