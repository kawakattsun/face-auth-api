ARG GOLANG_VERSION=1.12.6

FROM golang:${GOLANG_VERSION}-alpine

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
