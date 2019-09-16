# Use go 1.x based on alpine image.
FROM golang:1.13.0-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY directory-inspection-api   /go/src/go.nlx.io/nlx/directory-inspection-api
COPY directory-db/dbversion     /go/src/go.nlx.io/nlx/directory-db/dbversion
COPY common                     /go/src/go.nlx.io/nlx/common
COPY .git                       /go/src/go.nlx.io/nlx/directory-inspection-api/.git

WORKDIR /go/src/go.nlx.io/nlx/directory-inspection-api

RUN GIT_COMMIT_HASH=$(git rev-parse HEAD) && \
        GIT_TAG_NAME=$(git describe --tags) && \
        go build \
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
