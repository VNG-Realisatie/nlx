# Use go 1.x based on alpine image.
FROM golang:1.18.4-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building inway
COPY inway                                          /go/src/go.nlx.io/nlx/inway
COPY directory-monitor/health                       /go/src/go.nlx.io/nlx/directory-monitor/health
COPY directory-api/api                              /go/src/go.nlx.io/nlx/directory-api/api
COPY management-api                                 /go/src/go.nlx.io/nlx/management-api
COPY txlog-db/dbversion                             /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common                                         /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/inway

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME' " \
        -o dist/bin/nlx-inway ./cmd/nlx-inway

FROM alpine:3.16.0

RUN apk add --no-cache ca-certificates bash

COPY --from=build /go/src/go.nlx.io/nlx/inway/dist/bin/nlx-inway /usr/local/bin/nlx-inway

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-inway"]
