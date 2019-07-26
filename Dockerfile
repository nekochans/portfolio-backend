FROM golang:1.12.7-alpine3.10 as build

WORKDIR /go/app

COPY . .

RUN set -x && \
  apk update && \
  apk add --no-cache git && \
  go build -o portfolio-backend

FROM alpine:3.10

WORKDIR /app

COPY --from=build /go/app/portfolio-backend .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/portfolio-backend

CMD ["./portfolio-backend"]
