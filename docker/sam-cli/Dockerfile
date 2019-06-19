ARG GOLANG_VERSION=1.12.6
ARG AWS_CLI_VERSION=1.16.180
ARG AWS_SAM_CLI_VERSION=0.17.0

FROM golang:${GOLANG_VERSION}-alpine

ENV AWS_CLI_VERSION ${AWS_CLI_VERSION}
ENV AWS_SAM_CLI_VERSION ${AWS_SAM_CLI_VERSION}

ADD packages_requirements.txt ./

RUN apk add --no-cache \
        python3 \
        python3-dev \
        py-pip \
        git \
        make \
        bash \
        gcc \
        libc-dev \
        openssl \
        curl \
    && \
    pip3 install --upgrade \
        pip \
    && \
    pip3 install --user -r packages_requirements.txt