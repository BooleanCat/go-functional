.PHONY: test test-unit test-integration install

test:
	./scripts/test.sh

test-unit:
	./scripts/test.sh --regexScansFilePath --skip integration/

test-integration:
	./scripts/test.sh integration/

install:
	go install github.com/BooleanCat/go-functional
