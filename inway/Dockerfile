# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git openssh-client gcc musl-dev ca-certificates

# Install modd, which is used within docker-compose.
RUN go get github.com/cortesi/modd/cmd/modd

# Only add code that we use for building inway
COPY inway                      /go/src/go.nlx.io/nlx/inway
COPY directory-monitor/health   /go/src/go.nlx.io/nlx/directory-monitor/health
COPY directory/directoryapi     /go/src/go.nlx.io/nlx/directory/directoryapi
COPY txlog-db/dbversion         /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common                     /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/inway

# Ignore failing make to allow this stage to complete even when there are errors in the sourcecode.
RUN make || true


# Release binary on latest alpine image.
FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

COPY --from=build /go/src/go.nlx.io/nlx/inway/dist/bin/nlx-inway /usr/local/bin/nlx-inway


# TODO: shouldn't be in container, instead move to secret/map, or wait until config file is removed at all (replaced by central managing interface)
COPY --from=build /go/src/go.nlx.io/nlx/inway/service-config.toml /service-config.toml
ENV SERVICE_CONFIG /service-config.toml

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

EXPOSE 2018
CMD ["/usr/local/bin/nlx-inway"]
