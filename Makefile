.PHONY: test lint vet test-unit test-integration generate-fixtures

ginkgo := go run github.com/onsi/ginkgo/ginkgo --race --randomizeAllSpecs -r
lint := go run github.com/golangci/golangci-lint/cmd/golangci-lint
go-functional := go run github.com/BooleanCat/go-functional

test: test-unit test-integration

vet:
	go vet ./...

test-unit: vet lint
	$(ginkgo) --skipPackage integration

test-integration: generate-fixtures
	$(ginkgo) integration/

lint:
	$(lint) run

generate-fixtures: clean-fixtures
	cd fixtures && $(go-functional) int
	cd fixtures && $(go-functional) string
	cd fixtures && $(go-functional) '*int'
	cd fixtures && $(go-functional) '*string'
	cd fixtures && $(go-functional) interface{}
	cd fixtures && $(go-functional) --import-path os FileMode
	cd fixtures && $(go-functional) --import-path os *File

clean-fixtures:
	rm -r fixtures/* || true
