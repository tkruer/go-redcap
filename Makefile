.PHONY: test clean

test:
	@echo "Running tests..."
	@go test -v ./tests/...

clean:
	@echo "Cleaning test cache..."
	@go clean -testcache
