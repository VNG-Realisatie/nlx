# Use go 1.x based on alpine image.
FROM golang:1.18.4-alpine AS build

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory-monitor
COPY directory-monitor      /go/src/go.nlx.io/nlx/directory-monitor
COPY common                 /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/directory-monitor

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-directory-monitor ./cmd/nlx-directory-monitor

FROM alpine:3.16.0

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/go.nlx.io/nlx/directory-monitor/dist/bin/nlx-directory-monitor /usr/local/bin/nlx-directory-monitor

# Add non-priveleged user
RUN adduser -D -u 1001 appuser
USER appuser

CMD ["/usr/local/bin/nlx-directory-monitor"]
