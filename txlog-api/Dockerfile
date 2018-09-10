# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install go-based dependencies and tools
# RUN go get \
        # No extra pre-fetch packages currently required for txlog-api


# Only add code that we use for building directory
COPY txlog-api					/go/src/go.nlx.io/nlx/txlog-api
COPY txlog-db/dbversion			/go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common						/go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/txlog-api

# Ignore failing make to allow this stage to complete even when there are errors in the sourcecode.
RUN make || true 


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/txlog-api/dist/bin/nlx-txlog-api /usr/local/bin/nlx-txlog-api
