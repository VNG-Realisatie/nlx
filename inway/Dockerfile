# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git openssh-client gcc musl-dev ca-certificates

# Install modd, which is used within docker-compose.
RUN go get github.com/cortesi/modd/cmd/modd

# Add code.
COPY . /go/src/go.nlx.io/nlx
WORKDIR /go/src/go.nlx.io/nlx/inway

# Ignore failing make to allow this stage to complete even when there are errors in the sourcecode.
RUN make || true


# Release binary on latest alpine image.
FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/go.nlx.io/nlx/inway/dist/bin/nlx-inway /usr/local/bin/nlx-inway

EXPOSE 2018
CMD ["/usr/local/bin/nlx-inway"]
