DB_USER          := ${DB_USER}
DB_PASSWORD      := ${DB_PASSWORD}
DB_NAME          := ${DB_NAME}
DB_HOST          := ${DB_HOST}
TEST_DB_USER     := ${TEST_DB_USER}
TEST_DB_PASSWORD := ${TEST_DB_PASSWORD}
TEST_DB_NAME     := ${TEST_DB_NAME}

.PHONY: migrate-up migrate-down lint test
migrate-up:
	migrate -source file://./_sql -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:3306)/${DB_NAME}" up
	migrate -source file://./_sql -database "mysql://${TEST_DB_USER}:${TEST_DB_PASSWORD}@tcp(${DB_HOST}:3306)/${TEST_DB_NAME}" up
migrate-down:
	@migrate -source file://./_sql -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:3306)/${DB_NAME}" down
	@migrate -source file://./_sql -database "mysql://${TEST_DB_USER}:${TEST_DB_PASSWORD}@tcp(${DB_HOST}:3306)/${TEST_DB_NAME}" down
lint:
	go vet ./...
	gofmt -l -s -w .
	golangci-lint run ./...
test:
	go test -v ./...
