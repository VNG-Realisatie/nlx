# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install go-based dependencies and tools
RUN go get \
        github.com/cortesi/modd/cmd/modd

# Add code.
COPY . /go/src/go.nlx.io/nlx
WORKDIR /go/src/go.nlx.io/nlx/logdb-api

# Ignore failing make to allow this stage to complete even when there are errors in the sourcecode.
RUN make || true 


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/logdb-api/dist/bin/nlx-logdb-api /usr/local/bin/nlx-logdb-api
