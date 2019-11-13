FROM python:3.7.3-alpine3.9

ENV PYTHONUSERBASE=/usr/local

COPY packages_requirements.txt ./

RUN set -ex \
  && apk add --no-cache \
    bash \
    gcc \
    make \
    musl-dev \
    jq

RUN set -ex \
  && pip install --user -r packages_requirements.txt
