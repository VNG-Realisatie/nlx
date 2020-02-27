# Use go 1.x based on alpine image.
FROM golang:1.14.0-alpine AS build

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building config-api
COPY config-api                                     /go/src/go.nlx.io/nlx/config-api
COPY common                                         /go/src/go.nlx.io/nlx/common
COPY directory-registration-api                     /go/src/go.nlx.io/nlx/directory-registration-api

WORKDIR /go/src/go.nlx.io/nlx/config-api

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME' " \
        -o dist/bin/nlx-config-api ./cmd/nlx-config-api

# Release binary on latest alpine image.
FROM alpine:latest

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/config-api/dist/bin/nlx-config-api /usr/local/bin/nlx-config-api

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-config-api"]
