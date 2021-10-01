FROM golang:1.17-alpine3.14 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

ARG GOLANGCI_LINT_VERSION=v1.42.1
ARG AIR_VERSION=v1.27.3
ARG DLV_VERSION=v1.7.1
ARG OAPI_CODEGEN_VERSION=v1.8.2
ARG MIGRATE_VERSION=v4.15.0

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl make && \
  go install github.com/cosmtrek/air@${AIR_VERSION} && \
  go install github.com/go-delve/delve/cmd/dlv@${DLV_VERSION} && \
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION} && \
  go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@${MIGRATE_VERSION} && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION} && \
  go install golang.org/x/tools/cmd/goimports@latest

RUN set -eux && \
  go build -o portfolio-backend ./cmd/rest/main.go

ENV CGO_ENABLED 0

FROM alpine:3.14

WORKDIR /app

COPY --from=build /go/app/portfolio-backend .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/portfolio-backend

CMD ["./portfolio-backend"]
