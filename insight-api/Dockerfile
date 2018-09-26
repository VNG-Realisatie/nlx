# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install go-based dependencies and tools
# RUN go get \
        # No extra pre-fetch packages currently required for insight-api


# Only add code that we use for building directory
COPY insight-api        /go/src/go.nlx.io/nlx/insight-api
COPY txlog-db/dbversion /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common             /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/insight-api

RUN make


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/insight-api/dist/bin/nlx-insight-api /usr/local/bin/nlx-insight-api

# TODO: shouldn't be in container, instead move to secret/map, or wait until config file is removed at all (replaced by central managing interface)
COPY --from=build /go/src/go.nlx.io/nlx/insight-api/service-config.toml /service-config.toml
ENV SERVICE_CONFIG /service-config.toml
