# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory-monitor
COPY directory-monitor      /go/src/go.nlx.io/nlx/directory-monitor
COPY directory-db/dbversion /go/src/go.nlx.io/nlx/directory-db/dbversion
COPY common                 /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/directory-monitor

RUN go build -o dist/bin/nlx-directory-monitor ./cmd/nlx-directory-monitor


# Release binary on latest alpine image.
FROM alpine:latest

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-monitor/dist/bin/nlx-directory-monitor /usr/local/bin/nlx-directory-monitor
