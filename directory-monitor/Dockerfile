# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git protobuf

# Install go-based dependencies and tools
RUN go get \
        github.com/cortesi/modd/cmd/modd \
        google.golang.org/grpc \
        github.com/golang/protobuf/protoc-gen-go \
        github.com/gogo/protobuf/protoc-gen-gogofast \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# Only add code that we use for building directory-monitor
COPY directory-monitor      /go/src/go.nlx.io/nlx/directory-monitor
COPY directory-db/dbversion /go/src/go.nlx.io/nlx/directory-db/dbversion
COPY common                 /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/directory-monitor

RUN make


# Release binary on latest alpine image.
FROM alpine:latest

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-monitor/dist/bin/nlx-directory-monitor /usr/local/bin/nlx-directory-monitor
