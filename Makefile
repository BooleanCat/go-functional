.PHONY: test vet test-unit test-integration install generate-fixtures

ginkgo := ginkgo --race --randomizeAllSpecs -r

test: vet test-unit test-integration

vet:
	go vet ./gen/...
	go vet ./template/...
	go vet ./pkgname/...
	go vet ./

test-unit:
	$(ginkgo) gen/ template/ pkgname/

test-integration: install generate-fixtures
	$(ginkgo) integration/

install:
	go install github.com/BooleanCat/go-functional

generate-fixtures: install clean-fixtures
	cd fixtures && go-functional int
	cd fixtures && go-functional string
	cd fixtures && go-functional '*int'
	cd fixtures && go-functional '*string'
	cd fixtures && go-functional interface{}

clean-fixtures:
	rm -r fixtures/* || true
