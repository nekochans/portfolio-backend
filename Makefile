.PHONY: lint test
lint:
	@go vet ./...
	@gofmt -l -s -w .
test:
	@go test -v ./...
