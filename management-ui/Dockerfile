FROM node:14.17.4-alpine AS node-build

RUN apk add --no-cache --update git jq
RUN apk add --no-cache python3 make g++

ENV CI=true

# Only copy package.json and package-lock.json so we can use Docker's caching mechanism.
COPY management-ui/package.json \
    management-ui/package-lock.json \
    /go/src/go.nlx.io/nlx/management-ui/

WORKDIR /go/src/go.nlx.io/nlx/management-ui

RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY management-ui /go/src/go.nlx.io/nlx/management-ui

RUN npm run build

# Use go 1.x based on alpine image.
FROM golang:1.16.6-alpine AS go-build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY management-ui  /go/src/go.nlx.io/nlx/management-ui
COPY common  /go/src/go.nlx.io/nlx/common

WORKDIR /go/src/go.nlx.io/nlx/management-ui

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-management-ui ./cmd/nlx-management-ui

# Release binary on latest alpine image.
FROM alpine:3.14.0

COPY --from=go-build /go/src/go.nlx.io/nlx/management-ui/dist/bin/nlx-management-ui /usr/local/bin/nlx-management-ui
COPY --from=node-build /go/src/go.nlx.io/nlx/management-ui/build /app/public

WORKDIR /app

# Add non-privileged user
RUN adduser -D -u 1001 appuser && \
    chown -R 1001 /app

USER appuser

CMD ["/usr/local/bin/nlx-management-ui"]
