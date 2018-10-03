# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY txlog-api          /go/src/go.nlx.io/nlx/txlog-api
COPY txlog-db/dbversion /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common             /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/txlog-api

RUN go build -o dist/bin/nlx-txlog-api ./cmd/nlx-txlog-api


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/txlog-api/dist/bin/nlx-txlog-api /usr/local/bin/nlx-txlog-api
