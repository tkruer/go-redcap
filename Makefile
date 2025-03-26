.PHONY: test clean

test:
	@echo "Running tests..."
	@go test -v ./tests/...

test-coverage:
	@echo "Running tests with coverage..."
	@go test -cover ./...

clean:
	@echo "Cleaning test cache..."
	@go clean -testcache
