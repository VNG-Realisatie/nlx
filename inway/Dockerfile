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

# Only add code that we use for building inway
COPY inway                                          /go/src/go.nlx.io/nlx/inway
COPY directory-monitor/health                       /go/src/go.nlx.io/nlx/directory-monitor/health
COPY directory-registration-api/registrationapi     /go/src/go.nlx.io/nlx/directory-registration-api/registrationapi
COPY txlog-db/dbversion                             /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common                                         /go/src/go.nlx.io/nlx/common
COPY .git                                           /go/src/go.nlx.io/nlx/inway/.git

WORKDIR /go/src/go.nlx.io/nlx/inway

RUN GIT_COMMIT_HASH=$(git rev-parse HEAD) && \
        GIT_TAG_NAME=$(git describe --tags) && \
        go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME' " \
        -o dist/bin/nlx-inway ./cmd/nlx-inway

# Release binary on latest alpine image.
FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

COPY --from=build /go/src/go.nlx.io/nlx/inway/dist/bin/nlx-inway /usr/local/bin/nlx-inway


# TODO: 338 shouldn't be in container, instead move to secret/map, or wait until config file is removed at all (replaced by central managing interface)
COPY --from=build /go/src/go.nlx.io/nlx/inway/service-config.toml /service-config.toml
ENV SERVICE_CONFIG /service-config.toml

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-inway"]
