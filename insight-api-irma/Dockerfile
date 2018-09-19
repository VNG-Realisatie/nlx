# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install go-based dependencies and tools
# RUN go get \
        # No extra pre-fetch packages currently required for insight-api-irma


# Only add code that we use for building directory
COPY insight-api-irma   /go/src/go.nlx.io/nlx/insight-api-irma
COPY txlog-db/dbversion /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common             /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/insight-api-irma

RUN make


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/insight-api-irma/dist/bin/nlx-insight-api-irma /usr/local/bin/nlx-insight-api-irma
