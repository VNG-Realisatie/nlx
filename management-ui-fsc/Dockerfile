# Copyright © VNG Realisatie 2023
# Licensed under the EUPL

FROM node:20.1.0-alpine AS node-build

RUN apk add --update git python3 make g++

ENV CI=true

# Only copy package.json and package-lock.json so we can use Docker's caching mechanism.
COPY management-ui-fsc/ports/ui/assets/package.json \
    management-ui-fsc/ports/ui/assets/package-lock.json \
    /go/src/go.nlx.io/nlx/management-ui-fsc/

WORKDIR /go/src/go.nlx.io/nlx/management-ui-fsc

RUN npm ci --no-progress --color=false --quiet

# Now copy the whole workdir for the build step.
COPY management-ui-fsc/ports/ui/assets /go/src/go.nlx.io/nlx/management-ui-fsc

RUN npm run build

# Use go 1.x based on alpine image.
FROM golang:1.20.4-alpine AS go-build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Only add code that we use for building directory
COPY management-ui-fsc  /go/src/go.nlx.io/nlx/management-ui-fsc
COPY common             /go/src/go.nlx.io/nlx/common
COPY directory-api      /go/src/go.nlx.io/nlx/directory-api

WORKDIR /go/src/go.nlx.io/nlx/management-ui-fsc

ARG GIT_TAG_NAME=undefined
ARG GIT_COMMIT_HASH=undefined

RUN go build \
        -ldflags="-X 'go.nlx.io/nlx/common/version.BuildSourceHash=$GIT_COMMIT_HASH' -X 'go.nlx.io/nlx/common/version.BuildVersion=$GIT_TAG_NAME'" \
        -o dist/bin/nlx-management-ui-fsc ./

FROM alpine:3.17.3

COPY --from=go-build /go/src/go.nlx.io/nlx/management-ui-fsc/dist/bin/nlx-management-ui-fsc /usr/local/bin/nlx-management-ui-fsc
COPY --from=node-build /go/src/go.nlx.io/nlx/management-ui-fsc/dist /app/public/

WORKDIR /app

# Add non-privileged user
RUN adduser -D -u 1001 appuser && \
    chown -R 1001 /app

USER appuser

CMD ["/usr/local/bin/nlx-management-ui-fsc", "serve"]
