FROM golang:1.12.7-alpine3.10 as build

LABEL maintainer="https://github.com/nekochans"

ENV GO111MODULE on

WORKDIR /go/app

COPY . .

RUN set -ex && \
  apk update && \
  apk add --no-cache git && \
  go build -o portfolio-backend && \
  go get -u github.com/oxequa/realize && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine:3.10

WORKDIR /app

COPY --from=build /go/app/portfolio-backend .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/portfolio-backend

CMD ["./portfolio-backend"]
