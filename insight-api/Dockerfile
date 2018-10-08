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
COPY insight-api        /go/src/go.nlx.io/nlx/insight-api
COPY txlog-db/dbversion /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common             /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/insight-api

RUN go build -o dist/bin/nlx-insight-api ./cmd/nlx-insight-api


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/insight-api/dist/bin/nlx-insight-api /usr/local/bin/nlx-insight-api

# TODO:#344 shouldn't be in container, instead move to secret/map, or wait until config file is removed at all (replaced by central managing interface)
COPY --from=build /go/src/go.nlx.io/nlx/insight-api/insight-config.toml /insight-config.toml
ENV SERVICE_CONFIG /insight-config.toml
