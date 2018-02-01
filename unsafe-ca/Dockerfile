
# Build cfssl binaries
FROM golang:1-alpine AS build-cfssl
RUN apk add --update git g++
RUN go get -u github.com/cloudflare/cfssl/cmd/...

# Release image based on latest alpine.
FROM alpine:latest
# Add csfssl tools.
COPY --from=build-cfssl /go/bin/* /usr/local/bin/
# Add script to start ca server.
COPY unsafe-ca/*.sh /ca/
