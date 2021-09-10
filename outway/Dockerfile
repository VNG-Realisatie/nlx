# Use go 1.x based on alpine image.
FROM golang:1.17.1-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building outway
COPY outway                                     /go/src/go.nlx.io/nlx/outway
COPY directory-inspection-api/inspectionapi     /go/src/go.nlx.io/nlx/directory-inspection-api/inspectionapi
COPY txlog-db/dbversion                         /go/src/go.nlx.io/nlx/txlog-db/dbversion
COPY common                                     /go/src/go.nlx.io/nlx/common
COPY management-api                             /go/src/go.nlx.io/nlx/management-api

WORKDIR /go/src/go.nlx.io/nlx/outway

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME' " \
        -o  dist/bin/nlx-outway ./cmd/nlx-outway

FROM alpine:3.14.2
COPY --from=build /go/src/go.nlx.io/nlx/outway/dist/bin/nlx-outway /usr/local/bin/nlx-outway

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-outway"]
