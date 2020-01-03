FROM golang:1.13-alpine3.11 as build

LABEL maintainer="https://github.com/nekochans"

WORKDIR /go/app

COPY . .

RUN set -ex && \
  apk update && \
  apk add --no-cache git && \
  go build -o portfolio-backend && \
  go get gopkg.in/urfave/cli.v2@master && \
  go get github.com/oxequa/realize && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine:3.11

WORKDIR /app

COPY --from=build /go/app/portfolio-backend .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/portfolio-backend

CMD ["./portfolio-backend"]
