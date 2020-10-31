FROM golang:1.15-alpine3.12 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

ARG GOLANGCI_LINT_VERSION=v1.32.0

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl make && \
  go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv && \
  go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION} && \
  go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen && \
  go get golang.org/x/tools/cmd/goimports

ENV GO111MODULE on

RUN set -eux && \
  go build -o portfolio-backend ./cmd/rest/main.go

ENV CGO_ENABLED 0

FROM alpine:3.12

WORKDIR /app

COPY --from=build /go/app/portfolio-backend .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/portfolio-backend

CMD ["./portfolio-backend"]
