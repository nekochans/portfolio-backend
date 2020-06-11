.PHONY: migrate-up migrate-down lint test test-ci

migrate-up:
	@migrate -source file://./_sql -database 'mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):3306)/$(DB_NAME)' up
	@migrate -source file://./_sql -database 'mysql://$(TEST_DB_USER):$(TEST_DB_PASSWORD)@tcp($(DB_HOST):3306)/$(TEST_DB_NAME)' up
migrate-down:
	@migrate -source file://./_sql -database 'mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:3306)/${DB_NAME}' down
	@migrate -source file://./_sql -database 'mysql://${TEST_DB_USER}:${TEST_DB_PASSWORD}@tcp(${DB_HOST}:3306)/${TEST_DB_NAME}' down
lint:
	@go vet ./...
	@gofmt -l -s -w .
	@golangci-lint run ./...
test:
	@go test -v ./...
test-ci:
	@go test -v -coverprofile coverage.out -covermode atomic ./...
