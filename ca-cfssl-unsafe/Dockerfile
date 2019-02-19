FROM golang:1.11-alpine AS build-cfssl

# Build and install cfssl binaries
RUN apk add --update git g++
RUN go get -u github.com/cloudflare/cfssl/cmd/...

# Release image based on latest alpine.
FROM alpine:latest
# Add csfssl tools.
COPY --from=build-cfssl /go/bin/* /usr/local/bin/
# Add script to start ca server.
COPY *.sh /ca/

WORKDIR /ca
EXPOSE 8888
CMD ["./start-ca.sh", "nlx.local"]
