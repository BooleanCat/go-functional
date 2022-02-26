.phony: test check

check:
	@go vet ./...
	@gofmt -l .
	@test -z "$$( gofmt -l . )"

test:
	go test -v ./...
