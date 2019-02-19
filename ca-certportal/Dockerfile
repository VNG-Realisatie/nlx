# Use go 1.x based on the latest alpine image.
FROM golang:1.11-alpine AS build

# Install build tools.
RUN apk add --update git gcc musl-dev

# Cache dependencies
COPY go.mod /go/src/go.nlx.io/nlx/go.mod
COPY go.sum /go/src/go.nlx.io/nlx/go.sum
ENV GO111MODULE on
WORKDIR /go/src/go.nlx.io/nlx
RUN go mod download

# Add code and build.
COPY ca-certportal /go/src/go.nlx.io/nlx/ca-certportal
WORKDIR /go/src/go.nlx.io/nlx/ca-certportal
RUN go build -o dist/bin/nlx-ca-certportal ./cmd/nlx-ca-certportal


# Release binary on latest alpine image.
FROM alpine:latest

COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/dist/bin/nlx-ca-certportal /usr/local/bin/nlx-ca-certportal
COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/public /var/lib/nlx-ca-certportal/public

WORKDIR /var/lib/nlx-ca-certportal
CMD ["/usr/local/bin/nlx-ca-certportal"]
