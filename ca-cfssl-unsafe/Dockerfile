# Use go 1.x based on alpine image.
FROM golang:1.15.5-alpine AS build-cfssl

# Build and install cfssl binaries
RUN apk add --update git g++
RUN go get -u github.com/cloudflare/cfssl/cmd/...

# Release image based on latest alpine.
FROM alpine:latest

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
