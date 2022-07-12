# Use go 1.x based on alpine image.
FROM golang:1.18.4-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY directory-api  /go/src/go.nlx.io/nlx/directory-api
COPY txlog-api  /go/src/go.nlx.io/nlx/txlog-api
COPY management-api     /go/src/go.nlx.io/nlx/management-api
COPY common             /go/src/go.nlx.io/nlx/common
COPY outway             /go/src/go.nlx.io/nlx/outway

WORKDIR /go/src/go.nlx.io/nlx/management-api

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-management-api ./

FROM alpine:3.16.0
COPY --from=build /go/src/go.nlx.io/nlx/management-api/dist/bin/nlx-management-api /usr/local/bin/nlx-management-api

# Add non-privileged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-management-api", "serve"]
