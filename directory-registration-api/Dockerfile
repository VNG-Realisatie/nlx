# Use go 1.x based on alpine image.
FROM golang:1.16.1-alpine AS build

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building director
COPY directory-registration-api  /go/src/go.nlx.io/nlx/directory-registration-api
COPY directory-db/dbversion      /go/src/go.nlx.io/nlx/directory-db/dbversion
COPY common                      /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/directory-registration-api

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-directory-registration-api ./cmd/nlx-directory-registration-api

# Release binary on latest alpine image.
FROM alpine:latest

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-registration-api/dist/bin/nlx-directory-registration-api /usr/local/bin/nlx-directory-registration-api

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-directory-registration-api"]
