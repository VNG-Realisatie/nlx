# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install modd, which is used within docker-compose.
RUN go get github.com/cortesi/modd/cmd/modd

# Only add code that we use for building outway
COPY outway                 /go/src/go.nlx.io/nlx/outway
COPY directory/directoryapi /go/src/go.nlx.io/nlx/directory/directoryapi
COPY txlog-db/dbversion     /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common                 /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/outway

RUN make


# Release binary on latest alpine image.
FROM alpine:latest
COPY --from=build /go/src/go.nlx.io/nlx/outway/dist/bin/nlx-outway /usr/local/bin/nlx-outway

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

EXPOSE 12018
CMD ["/usr/local/bin/nlx-outway"]
