# Use go 1.x based on alpine image.
FROM golang:1.13.8-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY config-api         /go/src/go.nlx.io/nlx/config-api
COPY management-api     /go/src/go.nlx.io/nlx/management-api
COPY common             /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/management-api

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-management-api ./cmd/nlx-management-api

# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/management-api/dist/bin/nlx-management-api /usr/local/bin/nlx-management-api

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-management-api"]
