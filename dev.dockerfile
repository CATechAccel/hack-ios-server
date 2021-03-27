FROM golang:alpine3.12

WORKDIR /app

COPY ./ ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash curl mysql-client \
  && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && rm -rf /var/cache/apk/*
