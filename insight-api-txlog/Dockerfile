# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install go-based dependencies and tools
# RUN go get \
        # No extra pre-fetch packages currently required for insight-api-txlog


# Only add code that we use for building directory
COPY insight-api-txlog     /go/src/go.nlx.io/nlx/insight-api-txlog
COPY insight-api-irma/irma /go/src/go.nlx.io/nlx/insight-api-irma/irma
COPY txlog-db/dbversion    /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common                /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/insight-api-txlog

RUN make


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/insight-api-txlog/dist/bin/nlx-insight-api-txlog /usr/local/bin/nlx-insight-api-txlog
