.PHONY: migrate-up migrate-down lint format test ci

migrate-up:
	@migrate -source file://./_sql -database 'mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):3306)/$(DB_NAME)' up
	@migrate -source file://./_sql -database 'mysql://$(TEST_DB_USER):$(TEST_DB_PASSWORD)@tcp($(DB_HOST):3306)/$(TEST_DB_NAME)' up
migrate-down:
	@migrate -source file://./_sql -database 'mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:3306)/${DB_NAME}' down
	@migrate -source file://./_sql -database 'mysql://${TEST_DB_USER}:${TEST_DB_PASSWORD}@tcp(${DB_HOST}:3306)/${TEST_DB_NAME}' down
lint:
	@go vet ./...
	@golangci-lint run ./...
format:
	@gofmt -l -s -w .
	@goimports -w -l ./
test:
	@go clean -testcache
	@go test -p 1 -v ./...
ci: lint
	@go clean -testcache
	@go test -p 1 -v -coverprofile=covprofile.out -covermode atomic ./...
