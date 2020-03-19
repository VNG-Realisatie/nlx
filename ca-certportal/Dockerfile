# Use go 1.x based on alpine image.
FROM golang:1.14.1-alpine AS build

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Add code and build.
COPY ca-certportal      /go/src/go.nlx.io/nlx/ca-certportal
COPY common             /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/ca-certportal

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME' " \
        -o dist/bin/nlx-ca-certportal ./cmd/nlx-ca-certportal

# Release binary on latest alpine image.
FROM alpine:latest

COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/dist/bin/nlx-ca-certportal /usr/local/bin/nlx-ca-certportal
COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/public /var/lib/nlx-ca-certportal/public

WORKDIR /var/lib/nlx-ca-certportal

# Add non-priveleged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-ca-certportal"]
