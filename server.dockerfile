FROM golang:alpine3.12

WORKDIR /hack-ios-server

COPY ./ ./

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache bash curl mysql-client \
  && rm -rf /var/cache/apk/*
