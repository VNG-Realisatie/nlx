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
COPY directory-inspection-api   /go/src/go.nlx.io/nlx/directory-inspection-api
COPY directory-db/dbversion     /go/src/go.nlx.io/nlx/directory-db/dbversion
COPY common                     /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/directory-inspection-api
RUN go build -o dist/bin/nlx-directory-inspection-api ./cmd/nlx-directory-inspection-api


# Release binary on latest alpine image.
FROM alpine:latest

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-inspection-api/dist/bin/nlx-directory-inspection-api /usr/local/bin/nlx-directory-inspection-api
