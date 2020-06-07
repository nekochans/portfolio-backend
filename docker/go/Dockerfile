FROM golang:1.14-alpine3.11 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

ENV GO111MODULE=off

ARG GOLANGCI_LINT_VERSION=v1.27.0

RUN set -eux && \
  apk update && \
  apk add --no-cache git curl make && \
  curl -fLo /go/bin/air https://git.io/linux_air && \
  chmod +x /go/bin/air && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv && \
  go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}

ENV GO111MODULE on

RUN set -eux && \
  go build -o portfolio-backend ./cmd/rest/main.go

ENV CGO_ENABLED 0

ENV GCP_PROJECT_ID ${GCP_PROJECT_ID}
ENV DB_USER ${DB_USER}
ENV DB_PASSWORD ${DB_PASSWORD}
ENV DB_NAME ${DB_NAME}
ENV DB_HOST ${DB_HOST}
ENV TEST_DB_USER ${TEST_DB_USER}
ENV TEST_DB_PASSWORD ${TEST_DB_PASSWORD}
ENV TEST_DB_NAME ${TEST_DB_NAME}

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/portfolio-backend .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/portfolio-backend

CMD ["./portfolio-backend"]
