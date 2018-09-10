# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install modd, which is used within docker-compose.
RUN go get github.com/cortesi/modd/cmd/modd

# Add code and build.
COPY . /go/src/go.nlx.io/nlx/ca-certportal
WORKDIR /go/src/go.nlx.io/nlx/ca-certportal
RUN make


# Release binary on latest alpine image.
FROM alpine:latest

COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/dist/bin/nlx-ca-certportal /usr/local/bin/nlx-ca-certportal
COPY --from=build /go/src/go.nlx.io/nlx/ca-certportal/files /home/nlx-ca-certportal/files

WORKDIR /home/nlx-ca-certportal
EXPOSE 12020
CMD ["/usr/local/bin/nlx-ca-certportal"]
