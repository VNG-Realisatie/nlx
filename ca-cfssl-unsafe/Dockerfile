# Use go 1.x based on alpine image.
FROM golang:1.17.1-alpine AS build-cfssl

# Disable go-modules
ENV GO111MODULE=off

# Build and install cfssl binaries
RUN apk add --update git g++
RUN go get -u github.com/cloudflare/cfssl/cmd/...

FROM alpine:3.14.2

# Add csfssl tools.
COPY --from=build-cfssl /go/bin/* /usr/local/bin/

# Add non-privileged user
RUN adduser -D -u 1001 appuser

# Add script to start ca server.
COPY *.sh /ca/
RUN chown 1001:1001 /ca

WORKDIR /ca
EXPOSE 8888

USER appuser

CMD ["./start-ca.sh", "nlx.local"]
