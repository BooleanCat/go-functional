.PHONY: test test-unit test-integration install generate-fixtures

ginkgo := ginkgo --race --randomizeAllSpecs -r

test: test-unit test-integration

test-unit:
	$(ginkgo) gen/ template/

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
