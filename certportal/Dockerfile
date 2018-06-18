# Use go 1.x based on the latest alpine image.
FROM golang:1-alpine AS build

# Install build tools.
RUN apk add --update make git

# Install modd, which is used within docker-compose.
RUN go get github.com/cortesi/modd/cmd/modd

# Add code and build.
COPY . /go/src/go.nlx.io/nlx
WORKDIR /go/src/go.nlx.io/nlx/certportal
RUN make


# Release binary on latest alpine image.
FROM alpine:latest

COPY --from=build /go/src/go.nlx.io/nlx/certportal/dist/bin/nlx-certportal /usr/local/bin/nlx-certportal
COPY --from=build /go/src/go.nlx.io/nlx/certportal/files /home/nlx-certportal/files

WORKDIR /home/nlx-certportal
EXPOSE 12020
CMD ["/usr/local/bin/nlx-certportal"]
