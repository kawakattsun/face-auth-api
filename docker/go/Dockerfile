FROM golang:1.12.6-alpine3.9

RUN apk add --no-cache \
        git \
        make \
        bash \
        gcc \
        libc-dev \
        openssl \
        curl \
    && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u golang.org/x/lint/golint

RUN dep ensure
