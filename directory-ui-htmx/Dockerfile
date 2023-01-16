# Copyright © VNG Realisatie 2023
# Licensed under the EUPL

## Dart stage
FROM dart:stable AS static

COPY directory-ui-htmx /directory-ui-htmx

WORKDIR /dart-sass
RUN git clone https://github.com/sass/dart-sass.git . && \
  dart pub get && \
  dart ./bin/sass.dart --style=compressed \
    /directory-ui-htmx/ports/ui/scss/main.scss \
    /directory-ui-htmx/ports/ui/static/styling/main.css

# Use go 1.x based on alpine image.
FROM golang:1.20.2-alpine AS go-build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY directory-ui-htmx  /go/src/go.nlx.io/nlx/directory-ui-htmx
COPY common             /go/src/go.nlx.io/nlx/common
COPY directory-api      /go/src/go.nlx.io/nlx/directory-api

WORKDIR /go/src/go.nlx.io/nlx/directory-ui-htmx

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-directory-ui ./

FROM alpine:3.17.3

COPY --from=go-build /go/src/go.nlx.io/nlx/directory-ui-htmx/dist/bin/nlx-directory-ui /usr/local/bin/nlx-directory-ui
COPY --from=static /directory-ui-htmx/ports/ui/static /app/public/

WORKDIR /app

# Add non-privileged user
RUN adduser -D -u 1001 appuser && \
    chown -R 1001 /app

USER appuser

CMD ["/usr/local/bin/nlx-directory-ui", "serve"]
