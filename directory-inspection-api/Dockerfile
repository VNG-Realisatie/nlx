# Use go 1.x based on alpine image.
FROM golang:1.14.2-alpine AS build

ENV GO111MODULE on
ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
WORKDIR /go/src/go.nlx.io/nlx
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
RUN go mod download

# Only add code that we use for building directory
WORKDIR /go/src/go.nlx.io/nlx/directory-inspection-api
COPY common                     /go/src/go.nlx.io/nlx/common
COPY directory-db/dbversion     /go/src/go.nlx.io/nlx/directory-db/dbversion
COPY directory-inspection-api   /go/src/go.nlx.io/nlx/directory-inspection-api

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-directory-inspection-api ./cmd/nlx-directory-inspection-api

# Release binary on latest alpine image.
FROM alpine:latest

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-inspection-api/dist/bin/nlx-directory-inspection-api /usr/local/bin/nlx-directory-inspection-api

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-directory-inspection-api"]
